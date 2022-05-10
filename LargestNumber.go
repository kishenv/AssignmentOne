
///Problem statement : Generate the largest possible number from a given array of numbers.
// Input : [9,93,24,6]
// Output : [9,93,6,24]

package main

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	numArray            []string
	largestNum          []string
	userInput           string
	containsNonZeroNums bool
	inputLength         int
)

// largestNumGenerator
// Function to generate the largest number from the inputs provided by user.
func largestNumGenerator(numArray []string) {
	containsNonZeroNums = false
	inputLength = len(numArray)
	for i := 0; i < inputLength; i++ {
		for j := i + 1; j < inputLength; j++ {
			if numArray[i]+numArray[j] < numArray[j]+numArray[i] {
				numArray[i], numArray[j] = numArray[j], numArray[i]

			}
		}
	}
	for _, i := range numArray {
		val, ok := strconv.Atoi(i)
		if ok != nil {
			fmt.Print("ERROR: Atoi conversion failed! Input contains non integer values/ bad input format")
			return
		}
		if val != 0 {
			containsNonZeroNums = true
		}
		largestNum = append(largestNum, i)
	}
	if !containsNonZeroNums {
		fmt.Print([]string{"0"}) // Corner case where the input is a series of zeroes.
	} else {
		fmt.Print("[")
		for ind, i := range largestNum {
			if inputLength-1 == ind {
				fmt.Print(i, "]")
			} else {
				fmt.Print(i, ",")
			}
		}
	}
}

func main() {
	fmt.Println("Enter an array of numbers to generate the largest possible number. E.g. [3,4,5]")
	fmt.Scanf("%s", &userInput)
	userInput = userInput[1 : len(userInput)-1]
	numArray = strings.Split(userInput, ",")
	largestNumGenerator(numArray)
}
