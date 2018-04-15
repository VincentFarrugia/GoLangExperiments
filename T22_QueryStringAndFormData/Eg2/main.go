package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("Templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", rootEndpointHandleFunc)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("WebContent"))))
	http.ListenAndServe(":8080", nil)
}

func rootEndpointHandleFunc(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "homepage.gohtml", req.FormValue("txtBxA.Text"))
}
