// Client for this example.

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const cCmdMsgGoodbye = "<--goodbye-->"

var bServerTerminatedConnection = false
var serverConn net.Conn
var exitChan chan bool
var numMsgReceived int
var bReadingMyMsg = false

func main() {
	fmt.Println("*******************************************")
	fmt.Println("I am a Client! Attempting to connect to server.")
	fmt.Println("*******************************************")
	serverConn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Could not connect to Server. Exiting.")
		return
	}
	exitChan = make(chan bool)
	go handleMessagesFromServer(serverConn, exitChan)
	go handleMessagesToServer(serverConn)
	<-exitChan
	fmt.Println("Client terminated.")
}

//////////////////////////////////////////
// GOROUTINES:
//////////////////////////////////////////

func handleMessagesFromServer(conn net.Conn, exitChan chan bool) {
	servScanner := bufio.NewScanner(conn)
	for (!bServerTerminatedConnection) && (servScanner.Scan()) {
		receivedText := servScanner.Text()
		bServerTerminatedConnection = (receivedText == cCmdMsgGoodbye)
		if bServerTerminatedConnection {
			displayMessage("ComProt", "Received server connection close. Now ending client.")
			conn.Close()
			break
		} else {
			displayMessage("Server", servScanner.Text())
		}
		numMsgReceived++
	}
	fmt.Println("go routine handleMessagesFromServer terminated.")
	exitChan <- true
}

func handleMessagesToServer(conn net.Conn) {

	for numMsgReceived <= 0 {
		// Wait until we get our first message before
		// enabling the user stdin.
	}

	stdInScanner := bufio.NewScanner(os.Stdin)
	for !bServerTerminatedConnection {
		fmt.Print(">: ")
		bReadingMyMsg = true
		stdInScanner.Scan()
		sendMessage(conn, stdInScanner.Text())
	}
	bReadingMyMsg = false
	fmt.Println("go routine handleMessagesToServer terminated.")
}

//////////////////////////////////////////
// HELPER FUNCTIONS:
//////////////////////////////////////////

func sendMessage(conn net.Conn, msg string) {
	//fmt.Printf("I said: '%s'\n", msg)
	fmt.Fprintf(conn, "%s\n", msg)
}

func displayMessage(speakerName, msg string) {
	if bReadingMyMsg {
		fmt.Println()
	}
	fmt.Printf("%s said: %s\n", speakerName, msg)
	if bReadingMyMsg && !bServerTerminatedConnection {
		fmt.Print(">: ")
	}
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

//////////////////////////////////////////
