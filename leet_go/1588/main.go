package main

import (
	"fmt"
	"math"
)

func sumOddLengthSubarrays(arr []int) int {

	currentSub := 1
	totalLoop := int(math.Ceil(float64(len(arr) / 2)))
	total := 0
	for i := 0; i <= totalLoop; i++ {
		fmt.Println("======================")
		fmt.Println("SUB: ", currentSub)
		for j := 0; j+currentSub <= len(arr); j++ {
			fmt.Println("SUB Value: ", arr[j:j+currentSub])
			for _, num := range arr[j : j+currentSub] {
				fmt.Println("ADD: ", num)
				total += num
			}
		}
		currentSub += 2

	}

	return total
}

func main() {
	fmt.Println("Result: ", sumOddLengthSubarrays([]int{1, 4, 2, 5, 3}))
}
