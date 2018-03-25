package main

import (
	"html/template"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("Templates/*.gohtml"))
}

func main() {

	fd, err := os.Create("myTestData.html")
	panicIfErr(err)
	defer fd.Close()

	name := "Censu"
	err = tpl.ExecuteTemplate(fd, "tplDataTest.gohtml", name)
	panicIfErr(err)
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
