package main

import (
	"fmt"
)

func main() {
	generateMatrix(5)
}

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	// Print (initial state) row by row
	fmt.Println("Initial:")
	for i := 0; i < n; i++ {
		fmt.Println(res[i])
	}

	top, bottom := 0, n-1
	left, right := 0, n-1
	num := 1

	// Go "while" is just for + condition
	for top <= bottom && left <= right {

		// left -> right (top row)
		for c := left; c <= right; c++ {
			res[top][c] = num
			num++
		}
		top++

		// top -> bottom (right col)
		for r := top; r <= bottom; r++ {
			res[r][right] = num
			num++
		}
		right--

		// right -> left (bottom row) (only if still valid)
		if top <= bottom {
			for c := right; c >= left; c-- {
				res[bottom][c] = num
				num++
			}
			bottom--
		}

		// bottom -> top (left col) (only if still valid)
		if left <= right {
			for r := bottom; r >= top; r-- {
				res[r][left] = num
				num++
			}
			left++
		}
	}

	// Print (after fill) row by row
	fmt.Println("After fill:")
	for i := 0; i < n; i++ {
		fmt.Println(res[i])
	}

	return res
}
