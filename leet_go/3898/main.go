package main

import "fmt"

func findDegrees(matrix [][]int) []int {
	result := make([]int, len(matrix[0]))
	for i := 0; i < len(matrix[0]); i++ {
		for _, val := range matrix[i] {
			result[i] += val
		}
	}

	return result
}

func main() {
	fmt.Println("findDegrees: ", findDegrees([][]int{{0, 1, 1}, {1, 0, 1}, {1, 1, 0}}))
}
