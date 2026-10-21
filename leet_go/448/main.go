package main

import (
	"fmt"
)

func main() {
	fmt.Println("result: ", findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}

// {4, 3, 2, 7, 8, 2, 3, 1}
// {1, 2, 3, 4, 5, 6, 7, 8}

// {1, 2, 3, -4, 5, 6, 7, 8}

func findDisappearedNumbers(nums []int) []int {
	// Step 1: Mark visited indices by negating the value
	for _, num := range nums {
		idx := abs(num) - 1 // convert value to index
		if nums[idx] > 0 {
			nums[idx] = -nums[idx]
		}
	}

	// Step 2: Collect indices where value is still positive
	var result []int
	for i, num := range nums {
		if num > 0 {
			result = append(result, i+1)
		}
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findDisappearedNumbers1(nums []int) []int {

	mapNum := make(map[int]struct{})

	for _, num := range nums {
		mapNum[num] = struct{}{}
	}
	var result []int
	for i := 1; i <= len(nums); i++ {
		if _, exist := mapNum[i]; !exist {
			result = append(result, i)
		}
	}

	return result
}
