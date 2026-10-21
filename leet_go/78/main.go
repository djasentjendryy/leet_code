package main

import "fmt"

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}

func subsets(nums []int) [][]int {
	var result [][]int
	return getCombinations(0, nums, []int{}, result)
}

func getCombinations(index int, nums []int, combination []int, result [][]int) [][]int {
	tmp := append([]int(nil), combination...)
	result = append(result, tmp)

	for i := index; i < len(nums); i++ {
		combination = append(combination, nums[i])
		result = getCombinations(i+1, nums, combination, result)
		combination = combination[:len(combination)-1]
	}

	return result
}
