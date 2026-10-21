package main

import "fmt"

func main() {
	fmt.Println("Result: ", intersection([]int{1, 2, 2, 1}, []int{2, 2}))
}

func intersection(nums1 []int, nums2 []int) []int {

	mapNum1 := make(map[int]bool)

	for _, num := range nums1 {
		mapNum1[num] = false
	}

	var result []int

	for _, num := range nums2 {
		if v, exist := mapNum1[num]; exist {
			if v {
				continue
			}
			result = append(result, num)
			mapNum1[num] = true
		}
	}

	return result
}
