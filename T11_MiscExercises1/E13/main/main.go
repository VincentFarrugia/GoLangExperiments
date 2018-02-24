///////////////////////////////////////////////////////////////////
// Exercise description:
// Project Euler.net
// Problem 42: Coded triangle numbers.
// https://projecteuler.net/problem=42
///////////////////////////////////////////////////////////////////

package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"unicode"
)

func main() {
	runFileVersion()
	//runInteractiveVersion()
}

func runFileVersion() {

	csvFile, fileOpenErr := os.Open("p042_words.txt")
	if fileOpenErr != nil {
		fmt.Println("An Error has occurred while trying to open the CSV file")
		return
	}
	csvFileReader := csv.NewReader(csvFile)
	records, err := csvFileReader.ReadAll()
	if err != nil {
		fmt.Println("An Error has occurred while parsing the inputted CSV file")
	} else {
		numTriangleWordsInFile := 0
		numWords := len(records[0])
		for i := 0; i < numWords; i++ {
			bFlag, _ := isWordTriangleWord(records[0][i])

			if bFlag {
				numTriangleWordsInFile++
			}
		}

		fmt.Println("Number of Triangle Words in CSV file is: ", numTriangleWordsInFile)
	}
}

func runInteractiveVersion() {
	inWord := "HELLO"
	fmt.Print("Enter an English word: ")
	fmt.Scan(&inWord)

	bWordIsTriangleNumber, wordValue := isWordTriangleWord(inWord)

	fmt.Println("Word '", inWord, "' has a word value of ", wordValue)

	if bWordIsTriangleNumber {
		fmt.Println("Word '", inWord, "' is a Triangle Number")
	} else {
		fmt.Println("Word '", inWord, "' is NOT a Triangle Number")
	}
}

func isWordTriangleWord(word string) (bIsTriangleWord bool, wordValue int) {
	wordValue = calculateWordValue(word)
	return isTriangleNumber(wordValue), wordValue
}

func calculateWordValue(word string) int {
	wordValue := 0
	wordAsRuneSlice := []rune(word)
	totalNumChars := len(wordAsRuneSlice)
	for i := 0; i < totalNumChars; i++ {
		wordValue += getPositionOfLetterInAlphabet(wordAsRuneSlice[i])
	}
	return wordValue
}

func getPositionOfLetterInAlphabet(x rune) int {
	capitalisedRune := unicode.ToUpper(x)
	// A == 65
	// +1 so that the result is 1-based
	return int(capitalisedRune-65) + 1
}

func isTriangleNumber(x int) bool {

	idx := 1
	tRes := 0.0
	bBreakOut := false
	for !bBreakOut {
		tRes = 0.5 * (float64(idx) * float64(idx+1))
		if int(tRes) == x {
			return true
		}

		if tRes > float64(x) {
			return false
		}

		idx++
	}
	return false
}
