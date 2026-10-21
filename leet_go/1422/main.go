package main

import "fmt"

// Input: s = "011101"
// Output: 5
// Explanation:
// All possible ways of splitting s into two non-empty substrings are:
// left = "0" and right = "11101", score = 1 + 4 = 5
// left = "01" and right = "1101", score = 1 + 3 = 4
// left = "011" and right = "101", score = 1 + 2 = 3
// left = "0111" and right = "01", score = 1 + 1 = 2
// left = "01110" and right = "1", score = 2 + 1 = 3

func maxScore1(s string) int {
	total0 := 0
	total1 := 0
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i : i+1])
		if s[i:i+1] == "0" {
			total0 += 1
		} else {
			total1 += 1
		}
	}
	fmt.Println("---------------------------------------------")
	result := 0
	for i := 1; i < len(s); i++ {
		temp0 := total0
		temp1 := total1
		currLeft := s[0:i]
		currRight := s[i:len(s)]
		fmt.Println("currLeft: ", currLeft, ", currRight: ", currRight)
		for j := 0; j < len(currLeft); j++ {
			currChar := currLeft[j : j+1]
			if currChar == "0" {
				temp0 -= 1
			} else {
				temp1 -= 1
			}
		}
		fmt.Println("Value left: ", total0-temp0, ", Value right: ", temp1)
		if result < (total0-temp0)+temp1 {
			result = (total0 - temp0) + temp1
		}
		fmt.Println("---------------------------------------------")
	}

	return result
}

func maxScore(s string) int {
	total1 := 0
	for _, v := range s {
		if v == '1' {
			total1 += 1
		}
	}
	left0 := 0
	right1 := total1
	result := 0

	for _, v := range s {
		if v == '0' {
			left0++
		} else {
			right1--
		}
		if left0+right1 > result {
			result = left0 + right1
		}
	}

	return result
}
func main() {
	fmt.Println("Result: ", maxScore("00"))
}
