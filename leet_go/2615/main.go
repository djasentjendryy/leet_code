package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("result : ", distance([]int{1, 3, 1, 1, 2}))
}

/**
Input: nums = [1,3,1,1,2]
Output: [5,0,3,4,0]
Explanation:
When i = 0, nums[0] == nums[2] and nums[0] == nums[3]. Therefore, arr[0] = |0 - 2| + |0 - 3| = 5.
When i = 1, arr[1] = 0 because there is no other index with value 3.
When i = 2, nums[2] == nums[0] and nums[2] == nums[3]. Therefore, arr[2] = |2 - 0| + |2 - 3| = 3.
When i = 3, nums[3] == nums[0] and nums[3] == nums[2]. Therefore, arr[3] = |3 - 0| + |3 - 2| = 4.
When i = 4, arr[4] = 0 because there is no other index with value 2.
**/

func distance(nums []int) []int64 {
	// Group indices by value
	numMap := make(map[int][]int)
	for i, num := range nums {
		numMap[num] = append(numMap[num], i)
	}
	result := make([]int64, len(nums))

	for _, indexes := range numMap {
		k := len(indexes)

		// Build prefix sum of indices
		prefix := make([]int64, k+1)
		for i, idx := range indexes {
			prefix[i+1] = prefix[i] + int64(idx)
		}


		for j, idx := range indexes {
			leftCount := int64(j)
			rightCount := int64(k - j - 1)

			leftSum := prefix[j]                // sum of indices to the left
			rightSum := prefix[k] - prefix[j+1] // sum of indices to the right

			result[idx] = int64(idx)*leftCount - leftSum +
				rightSum - int64(idx)*rightCount
		}
	}

	return result
}

func distance1(nums []int) []int64 {
	numMap := make(map[int][]int)
	for i, num := range nums {
		numMap[num] = append(numMap[num], i)
	}

	var result []int64
	for i, num := range nums {
		indexes := numMap[num]
		var val int64
		for _, idx := range indexes {
			val += int64(math.Abs(float64(i - idx)))
		}
		result = append(result, val)
	}

	return result
}
