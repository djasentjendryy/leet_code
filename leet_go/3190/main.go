package main

import (
	"fmt"
)

func main() {
	fmt.Print(minimumOperations([]int{1, 2, 3, 4}))
}

func minimumOperations(nums []int) int {

	var result int = 0
	for _, num := range nums {
		remain := num % 3
		remainUpToNum := 3 - remain
		result += min(remain, remainUpToNum)
	}
	return result
}
