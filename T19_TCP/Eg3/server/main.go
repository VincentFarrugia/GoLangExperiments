// Server for this example.

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"time"
)

const cClientConnDeadline (time.Duration) = (10 * time.Second)
const cCmdMsgGoodbye = "<--goodbye-->"

var clientIDDispenser = -1

func main() {

	fmt.Println("*******************************************")
	fmt.Println("I am a SERVER! Let's see who connects!")
	fmt.Println("*******************************************")
	listener, err := net.Listen("tcp", "localhost:8080")
	defer listener.Close()
	panicIfErr(err)

	for {
		clientConn, err := listener.Accept()
		panicIfErr(err)
		if err != nil {
			clientConn.Close()
		}
		if err == nil && clientConn != nil {
			clientIDDispenser++
			go handleClient(clientConn, clientIDDispenser)
		}
	}
}

//////////////////////////////////////////
// GOROUTINES:
//////////////////////////////////////////

func handleClient(clientConn net.Conn, clientID int) {
	clientIDAsStr := fmt.Sprintf("C%d", clientID)
	fmt.Println()
	fmt.Printf("************ NEW CLIENT with ID: (%s) ************", clientIDAsStr)
	fmt.Println()
	sendMessage(clientConn, clientIDAsStr, fmt.Sprintf("Yo client what up? We got %d seconds so you gotta be quick yo!", cClientConnDeadline/time.Second))
	clientConn.SetReadDeadline(time.Now().Add(10 * time.Second))
	scanner := bufio.NewScanner(clientConn)
	responsesList := []string{"Hmm...", "Ok...", "Interesting...", "I see..."}
	for scanner.Scan() {
		clientText := scanner.Text()
		displayMessage(clientIDAsStr, clientText)
		if clientText == cCmdMsgGoodbye {
			break
		}
		sendMessage(clientConn, clientIDAsStr, responsesList[random(0, len(responsesList))])
	}
	sendMessage(clientConn, clientIDAsStr, "Thanks for stopping by! Talk to you later!")
	sendMessage(clientConn, clientIDAsStr, cCmdMsgGoodbye)
	clientConn.Close()
	fmt.Printf("************ CLIENT HANDLER (%s) DONE ************\n", clientIDAsStr)
	fmt.Println()
	fmt.Println()
}

//////////////////////////////////////////
// HELPER FUNCTIONS:
//////////////////////////////////////////

func sendMessage(conn net.Conn, clientIDAsStr string, msg string) {
	fmt.Printf("I said to (%s): '%s'\n", clientIDAsStr, msg)
	fmt.Fprintf(conn, "%s\n", msg)
}

func displayMessage(speakerName, msg string) {
	fmt.Printf("%s said: %s\n", speakerName, msg)
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return (min + rand.Intn(max-min))
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

//////////////////////////////////////////
