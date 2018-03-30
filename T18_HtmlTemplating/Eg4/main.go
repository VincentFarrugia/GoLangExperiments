package main

import (
	"html/template"
	"os"
	"strings"
)

var tpl *template.Template

var fnMap = template.FuncMap{
	"uc":  strings.ToUpper,
	"rev": reverseString,
}

type person struct {
	Fname string
	Lname string
}

func init() {
	tpl = template.Must(template.New("").Funcs(fnMap).ParseGlob("Templates/*.gohtml"))
}

func main() {

	personList := []person{
		{"Mark", "Borg"},
		{"David", "Spiteri"},
	}

	fd, err := os.Create("myTestData.html")
	panicIfErr(err)
	defer fd.Close()

	err = tpl.ExecuteTemplate(fd, "tplFunctionsTest.gohtml", personList)
	panicIfErr(err)
}

func reverseString(s string) string {
	s = strings.TrimSpace(s)
	strLen := len(s)
	halfLen := strLen / 2
	tmp := rune(0)
	sliceOfRune := []rune(s)
	for i := 0; i < halfLen; i++ {
		tmp = sliceOfRune[i]
		sliceOfRune[i] = sliceOfRune[strLen-1-i]
		sliceOfRune[strLen-1-i] = tmp
	}
	s = string(sliceOfRune)
	return s
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
