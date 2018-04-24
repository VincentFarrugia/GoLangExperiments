package main

import (
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var connSettings *DBConnSettings = nil
var dbConn DBConnection

func main() {

	// CHECK FOR VALID SETTINGS:
	if connSettings == nil {
		fmt.Println("No DB Conn Settings were provided.")
		return
	}

	// CREATE CONNECTION TO THE DATABASE:
	var connCreator DBConnCreator
	var err error
	dbConn, err = connCreator.CreateConnection(*connSettings)
	printOutErr(err)
	//defer dbConn.Close()

	// START THE HTTP SERVER:
	http.HandleFunc("/", rootEndpoint)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/showFriendsTable", showFriendsTableEndpoint)
	http.HandleFunc("/showBooksTable", showBooksTableEndpoint)
	http.HandleFunc("/createBooksTable", createBooksTableEndpoint)
	http.HandleFunc("/insertBook", insertBookEndpoint)
	http.HandleFunc("/dropBooksTable", dropBooksTableEndpoint)
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

func showFriendsTableEndpoint(w http.ResponseWriter, req *http.Request) {
	if !setResponseToErrorIfDBConnInvalid(w, req, dbConn) {
		rows, err := dbConn.RunQuery(`SELECT * FROM friends;`)
		printOutErr(err)
		defer rows.Close()
		retStr, err := convertRowsToString(rows)
		printOutErr(err)
		fmt.Fprintln(w, retStr)
	}
}

func showBooksTableEndpoint(w http.ResponseWriter, req *http.Request) {
	if !setResponseToErrorIfDBConnInvalid(w, req, dbConn) {
		rows, err := dbConn.RunQuery(`SELECT * FROM books;`)
		printOutErr(err)
		defer rows.Close()
		retStr, err := convertRowsToString(rows)
		printOutErr(err)
		fmt.Fprintln(w, retStr)
	}
}

func createBooksTableEndpoint(w http.ResponseWriter, req *http.Request) {
	if !setResponseToErrorIfDBConnInvalid(w, req, dbConn) {
		str := `CREATE TABLE books (name VARCHAR(20));`
		dbConn.RunMod(str)
		fmt.Fprintln(w, "Ran:'", str, "'")
	}
}

func insertBookEndpoint(w http.ResponseWriter, req *http.Request) {
	if !setResponseToErrorIfDBConnInvalid(w, req, dbConn) {
		str := `INSERT INTO books VALUES ("MobyDick");`
		dbConn.RunMod(str)
		fmt.Fprintln(w, "Ran:'", str, "'")
	}
}

func dropBooksTableEndpoint(w http.ResponseWriter, req *http.Request) {
	if !setResponseToErrorIfDBConnInvalid(w, req, dbConn) {
		str := `DROP TABLE books;`
		dbConn.RunMod(str)
		fmt.Fprintln(w, "Ran:'", str, "'")
	}
}

/////////////////////////////
// HELPER FUNCTIONS:
/////////////////////////////

func setResponseToErrorIfDBConnInvalid(w http.ResponseWriter, req *http.Request, conn DBConnection) bool {
	if IsValidDBConnectionNoPtr(conn) {
		return false
	}
	http.Error(w, "Invalid DB Connection", http.StatusInternalServerError)
	return true
}

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
