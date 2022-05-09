package main

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	numArray   []string
	largestNum string
	userInput  string
)

// largestNumGenerator
// Function to generate the largest number from the inputs provided by user.
func largestNumGenerator(numArray []string) string {
	for i := 0; i < len(numArray); i++ {
		for j := i + 1; j < len(numArray); j++ {
			if numArray[i]+numArray[j] > numArray[j]+numArray[i] {
				numArray[i], numArray[j] = numArray[j], numArray[i]

			}
		}
	}
	for _, i := range numArray {
		largestNum = i + largestNum
	}
	val, err := strconv.Atoi(largestNum)
	if err != nil {
		fmt.Print("ERROR: Atoi conversion failed! Input contains special-characters/alphabets")
		return ""
	}
	if val == 0 {
		return "0" // Corner case where the inputs provided is a series of 0s.
	} else {
		return largestNum
	}
}

func main() {
	fmt.Println("Enter a comma-separated list of numbers to generate largest number")
	fmt.Scanf("%s", &userInput)
	numArray = strings.Split(userInput, ",")
	print(largestNumGenerator(numArray))
}
