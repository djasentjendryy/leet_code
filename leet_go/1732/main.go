package main

import "fmt"

func largestAltitude(gain []int) int {

	result := 0
	temp := 0
	for _, v := range gain {
		temp += v
		if temp > result {
			result = temp
		}
	}

	return result
}

func main() {
	fmt.Println("Result: ", largestAltitude([]int{-4, -3, -2, -1, 4, 3, 2}))
}
