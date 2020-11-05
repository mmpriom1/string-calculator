package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	r, _ := Add("4,5,6")
	fmt.Println(r)
}

func Add(input string) (int, error) {
	if input == "" {
		return 0, nil
	}

	//considering if an input string starts with "//", it means it will have delimiter(s)
	if input[0:2] != "//" {
		result, err := addListOfNumbers(input)
		if err != nil {
			return -1, err
		}
		return result, nil
	}

	// handling arbitrary length delimiter

	delimiterList, startIndex := getMultipleDelimiter(input)
	numberStrings := input[startIndex:]

	// replacing all the delimiters with a common delimiter
	for _, delimiter := range delimiterList {
		numberStrings = strings.ReplaceAll(numberStrings, delimiter, ",")
	}

	result, err := addListOfNumbers(numberStrings)

	if err != nil {
		return -1, err
	}
	return result, nil
}

// getMultipleDelimiter will retrieve the list of delimiters of an input string. This should also handle delimiters with arbitrary length
func getMultipleDelimiter(input string) ([]string, int) {

	// splitting the string based on new line
	slices := strings.Split(input, "\n")

	// taking the first slice which will contain "//", delimiters and will end with a new line
	firstSlice := slices[0]

	// removing the "//" and new line
	deliMiterLists := firstSlice[2:(len(firstSlice))]

	// fetching all the delimiters separated by comma
	listOfDelimiter := strings.Split(deliMiterLists, ",")

	// startingIndex is the starting point of numbers with delimiters
	startingIndex := len(firstSlice) + 1

	return listOfDelimiter, startingIndex
}

// addListOfNumbers will add all the numbers from the input after removing all the delimiters
func addListOfNumbers(numberInput string) (int, error) {
	result := 0

	// replacing new line with comma
	replacedNewLine := strings.ReplaceAll(numberInput, "\n", ",")

	// replacing double comma with new line
	replacedNewLine = strings.ReplaceAll(replacedNewLine, ",,", ",")

	// getting the list of numbers
	listOfNumber := strings.Split(replacedNewLine, ",")

	// calculating result
	for _, number := range listOfNumber {
		r, _ := strconv.Atoi(number)
		if r < 0 {
			return -1, errors.New("negatives not allowed")
		}
		// Ignoring larger numbers
		if r > 1000 {
			continue
		}

		result += r
	}
	return result, nil
}
