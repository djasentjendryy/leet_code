package main

import (
	"fmt"
)

func minStartValue(nums []int) int {
	smallestVal := 0
	temp := 0
	for _, v := range nums {
		temp += v
		if temp < smallestVal {
			smallestVal = temp
		}
	}

	if smallestVal < 0 {
		return -smallestVal + 1
	}
	return 1
}

func main() {
	fmt.Println("Result: ", minStartValue([]int{-3, 2, -3, 4, 2}))
}
