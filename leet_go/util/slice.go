package util

func RemoveAll(nums []int, x int) []int {
	k := 0
	for _, v := range nums {
		if v != x {
			nums[k] = v
			k++
		}
	}
	return nums[:k]
}

func Clear(nums []int) []int {
	for idx, _ := range nums {
		nums[idx] = 0
	}
	return nums
}

func SumAll(nums []int) int {
	result := 0
	for _, v := range nums {
		result += v
	}
	return result
}
