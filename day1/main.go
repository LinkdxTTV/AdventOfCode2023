package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {

	file, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		testString := scanner.Text()
		fmt.Println(testString)

		stringArray := extractNumbersFromString(testString, true) // This should return []string{"1","5","3"}
		if len(stringArray) == 0 {
			fmt.Println("wtf this has no digits", testString)
			continue
		}

		twoDigitNumberAsString := stringArray[0] + stringArray[len(stringArray)-1]
		fmt.Println(stringArray)
		twoDigitNumber, err := strconv.Atoi(twoDigitNumberAsString)
		if err != nil {
			panic(err)
		}
		fmt.Println(twoDigitNumber)

		sum += twoDigitNumber
	}

	if err := scanner.Err(); err != nil && err != io.EOF {
		panic(err)
	}

	fmt.Println(sum)
}

var Numbers []string = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

type wasWord struct {
	number  string
	wasWord bool
}

func extractNumbersFromString(testString string, consideringWordsAsDigits bool) []string {

	stringArray := []string{}
	IndexNumbers := make([]wasWord, len(testString))

	for index, character := range testString {
		if unicode.IsDigit(character) {
			IndexNumbers[index] = wasWord{string(character), false}
		}
	}
	if consideringWordsAsDigits {
		for testNumberAsInteger, testNumberAsString := range Numbers {
			mutableString := testString
			var idx = 0
			var count = 0
			for len(mutableString) > 0 {
				idx = strings.Index(mutableString, testNumberAsString)
				if idx == -1 {
					break
				}
				IndexNumbers[idx+count] = wasWord{fmt.Sprint(testNumberAsInteger), true}
				mutableString = mutableString[idx+len(testNumberAsString):]
				count += idx + len(testNumberAsString)
			}
		}
	}

	for idx, numAsString := range IndexNumbers {
		if numAsString.number != "" {
			numAsInt, err := strconv.Atoi(numAsString.number)
			if err != nil {
				panic(err)
			}
			if numAsString.wasWord {
				for i := idx; i < idx+len(Numbers[numAsInt]); i++ {
					continue // Dont need.. twone should apparently be parsed 2, 1
					IndexNumbers[i].number = ""
					IndexNumbers[i].wasWord = false
				}
			}
			stringArray = append(stringArray, numAsString.number)
		}
	}
	return stringArray
}
