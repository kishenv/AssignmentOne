///Problem statement : Generate the largest possible number from a given array of numbers.
// Input :  9 93 24 6
// Output : [9 93 6 24]

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	numArray            []int
	largestNum          []int
	containsNonZeroNums bool
	inputLength         int
)

// largestNumGenerator
// Function to generate the largest number from an array of integers.
// Returns an array of integers arranged to generate the largest number
func largestNumGenerator(numArray []int) []int {
	containsNonZeroNums = false
	inputLength = len(numArray)
	for i := 0; i < inputLength; i++ {
		for j := i + 1; j < inputLength; j++ {
			toStri := strconv.Itoa(numArray[i])
			toStrj := strconv.Itoa(numArray[j])
			if toStri+toStrj < toStrj+toStri {
				numArray[i], numArray[j] = numArray[j], numArray[i]

			}
		}
	}
	for _, i := range numArray {
		if i != 0 {
			containsNonZeroNums = true
		}
		largestNum = append(largestNum, i)
	}
	if !containsNonZeroNums {
		return []int{0} // Corner case where the input is a series of zeroes.
	} else {
		return largestNum
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.Fields(scanner.Text())
	for _, i := range input {
		digit, err := strconv.Atoi(i)
		if err != nil {
			fmt.Print("Error: Non integer character present.")
			return
		}
		numArray = append(numArray, digit)
	}
	fmt.Println(largestNumGenerator(numArray))
}
