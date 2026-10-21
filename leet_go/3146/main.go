package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	findPermutationDifference("abcde", "edbac")
}

func findPermutationDifference(s string, t string) int {

	var result int
	for index, char := range s {
		currChar := string(char)
		//fmt.Println(currChar)
		index2 := strings.Index(t, currChar)
		//fmt.Println()

		result += int(math.Abs(float64(index) - float64(index2)))
	}

	fmt.Println(result)
	return 0
}
