package main

import (
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var connSettings *DBConnSettings = nil

func main() {

	// CHECK FOR VALID SETTINGS:
	if connSettings == nil {
		fmt.Println("No DB Conn Settings were provided.")
		return
	}

	// CREATE CONNECTION TO THE DATABASE:
	var connCreator DBConnCreator
	dbConn, err := connCreator.CreateConnection(*connSettings)
	printOutErr(err)
	defer dbConn.Close()

	// START THE HTTP SERVER:
	http.HandleFunc("/", rootEndpoint)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err = http.ListenAndServe(":8080", nil)
	printOutErr(err)
}

/////////////////////////////
// HTTP SERVER ENDPOINTS:
/////////////////////////////

func rootEndpoint(w http.ResponseWriter, req *http.Request) {
	_, err := io.WriteString(w, "Process Success!")
	printOutErr(err)
}

/////////////////////////////
// HELPER FUNCTIONS:
/////////////////////////////

func printOutErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

/////////////////////////////
