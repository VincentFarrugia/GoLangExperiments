package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

//////////////////////////////////////
// Golang Text/HTML templates 101.
// 1. Parse your template files.
// 2. Execute your template files with required data.
// 3. Save or send the resulting file.
//////////////////////////////////////

const tplHelloWorldFileName = "tplHelloWorld.gohtml"
const tplTestAFileName = "tplTestA.gohtml"
const tplTestBFileName = "tplTestB.gohtml"

func main() {
	deleteAnyOldHTMLFiles()
	testThree()
}

// TEST 1: Basic Parse and Execute
func testOne() {
	myTemplate, err := template.ParseFiles(tplHelloWorldFileName)
	handleError(err)

	outFD1, _ := createNwOutFile("index.html")
	defer outFD1.Close()

	err = myTemplate.Execute(outFD1, nil)
	handleError(err)
}

// TEST2: Parse in multiple template files and use ExecuteTemplate (i.e. not Execute)
// to select which template to use from our container of tempaltes.
func testTwo() {
	myOtherTemplates, err := template.ParseFiles(tplTestAFileName, tplTestBFileName)
	handleError(err)

	outFD2, _ := createNwOutFile("outTestA.html")
	defer outFD2.Close()
	outFD3, _ := createNwOutFile("outTestB.html")
	defer outFD3.Close()

	err = myOtherTemplates.ExecuteTemplate(outFD2, tplTestAFileName, nil)
	handleError(err)
	err = myOtherTemplates.ExecuteTemplate(outFD3, tplTestBFileName, nil)
	handleError(err)
}

// TEST3: Parsing in multiple template files using a Glob.
func testThree() {
	myTemplates, err := template.ParseGlob("Templates/*.gohtml")
	handleError(err)

	outFD4, _ := createNwOutFile("outTestC.html")
	defer outFD4.Close()
	outFD5, _ := createNwOutFile("outTestD.html")
	defer outFD5.Close()
	err = myTemplates.ExecuteTemplate(outFD4, "tplTestC.gohtml", nil)
	handleError(err)
	err = myTemplates.ExecuteTemplate(outFD5, "tplTestD.gohtml", nil)
}

/*// NOTE: To make things more performant you might want to load all your templates just once, at Init time.
var tpl *template.Template
func init() {
	// template.Must will perform basic error checking.
	tpl = template.Must(template.ParseGlob("Templates/*.gohtml"))
}*/

//////////////////////////////////////
// HELPER FUNCTIONS
//////////////////////////////////////

func createNwOutFile(path string) (*os.File, bool) {
	outFD, err := os.Create(path)
	handleErrorWithMsg("Could not create output file!", err)
	if err != nil && outFD != nil {
		outFD.Close()
	}
	return outFD, (err != nil)
}

func handleError(err error) {
	handleErrorWithMsg("", err)
}

func handleErrorWithMsg(programmerMsg string, err error) {
	if err != nil {
		if programmerMsg != "" {
			fmt.Println("Error: ", programmerMsg)
		}
		panic(err)
	}
}

func deleteAnyOldHTMLFiles() {
	currentDir := filepath.Dir(os.Args[0])
	deleteFilesByGlob(currentDir, "*.html")
	deleteFilesByGlob(filepath.Join(currentDir, "Templates"), "*.html")
}

func deleteFilesByGlob(dir string, globPattern string) {

	if (len(dir) <= 0) || (len(globPattern) <= 0) {
		return
	}

	matches, err := filepath.Glob(globPattern)
	handleError(err)
	for _, f := range matches {
		if err := os.Remove(f); err != nil {
			panic(err)
		}
	}
}

//////////////////////////////////////
