package main

import (
	"fmt"
)

// Input: nums = [1,7,3,6,5,6]
// Output: 3
// Explanation:
// The pivot index is 3.
// Left sum = nums[0] + nums[1] + nums[2] = 1 + 7 + 3 = 11
// Right sum = nums[4] + nums[5] = 5 + 6 = 11

// Input: nums = [1,2,3]
// Output: -1
// Explanation:
// There is no index that satisfies the conditions in the problem statement.

// Input: nums = [2,1,-1]
// Output: 0
// Explanation:
// The pivot index is 0.
// Left sum = 0 (no elements to the left of index 0)
// Right sum = nums[1] + nums[2] = 1 + -1 = 0

func pivotIndex1(nums []int) int {
	totalRight, totalLeft := 0, 0
	left, right := 0, len(nums)-1
	if nums[left] < nums[right] {
		totalLeft += nums[left]
	} else {
		totalRight += nums[right]
	}

	fmt.Println("left: ", left, " Right: ", right, "totalLeft: ", totalLeft, " totalRight: ", totalRight)

	for left < right {

		if totalLeft > totalRight {
			right--
			totalRight += nums[right]
		} else if totalLeft < totalRight {
			left++
			totalLeft += nums[left]
		} else {
			fmt.Println("DAPEEEET")
			fmt.Println("left: ", left, " Right: ", right, "totalLeft: ", totalLeft, " totalRight: ", totalRight)
			if left == 0 {
				return 0
			} else if right == len(nums) {
				return right
			} else {
				return left + 1
			}
		}
		fmt.Println("left: ", left, " Right: ", right, "totalLeft: ", totalLeft, " totalRight: ", totalRight)
	}
	return -1
}

func pivotIndex(nums []int) int {
	totalSum := 0
	for _, v := range nums {
		totalSum += v
	}

	leftSum := 0
	for i, v := range nums {
		if leftSum == totalSum-v-leftSum {
			return i
		}
		leftSum += v
	}

	return -1
}

func main() {
	fmt.Println("Result: ", pivotIndex([]int{1, 7, 3, 6, 5, 6}))
}
