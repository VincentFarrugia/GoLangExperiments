/*
Description:
Server which provides a page where the client can upload a file
which is then stored on the server machine.
*/

package main

import (
	"html/template"
	"io"
	"net/http"
	"os"
	"path"
)

var tpl *template.Template

const (
	cTemplatesDir       = "Templates"
	cWebContentDir      = "WebContent"
	cUserFileStorageDir = "UserFileStorage"
)

func init() {
	tpl = template.Must(template.ParseGlob(cTemplatesDir + "/*.gohtml"))
}

func main() {
	http.HandleFunc("/upload", uploadEndpointHandleFunc)
	http.HandleFunc("/", rootEndpointHandleFunc)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir(cWebContentDir))))
	http.ListenAndServe(":8080", nil)
}

func rootEndpointHandleFunc(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "homepage.gohtml", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func uploadEndpointHandleFunc(w http.ResponseWriter, req *http.Request) {

	if req.Method == http.MethodPost {
		reqFile, ptReqFileHd, err := req.FormFile("reqFile")
		if applyDefaultResponseIfError(w, err) {
			return
		}

		fd, err := os.Create(path.Join(cUserFileStorageDir, ptReqFileHd.Filename))
		if applyDefaultResponseIfError(w, err) {
			return
		}

		_, err = io.Copy(fd, reqFile)
		if applyDefaultResponseIfError(w, err) {
			return
		}
	}

	io.WriteString(w, "OK OK")
}

///////////////////////////////////
// HELPER FUNCTIONS:
///////////////////////////////////

func applyDefaultResponseIfError(w http.ResponseWriter, err error) bool {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return true
	}
	return false
}

///////////////////////////////////
