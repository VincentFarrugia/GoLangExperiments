package main

import (
	"os"
	"text/template"
)

var tpl *template.Template

type student struct {
	Fname string
	Lname string
}

type carManufacturer struct {
	Name    string
	Country string
}

type items struct {
	StudentList         []student
	CarManufacturerList []carManufacturer
}

func init() {
	tpl = template.Must(template.ParseGlob("Templates/*.gohtml"))
}

func main() {

	studentList := []student{
		{"David", "Farrugia"},
		{"Mark", "Borg"},
		{"Andrew", "Xerri"},
	}

	carManuList := []carManufacturer{
		{"Toyota", "Japan"},
		{"Landrover", "UK"},
		{"Fiat", "Italy"},
		{"Opel", "Germany"},
		{"Renault", "France"},
	}

	itemsInst := items{
		studentList,
		carManuList,
	}

	fd, err := os.Create("myTestData.html")
	panicIfErr(err)
	defer fd.Close()

	err = tpl.ExecuteTemplate(fd, "tplCompositeDataTest.gohtml", itemsInst)
	panicIfErr(err)
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
