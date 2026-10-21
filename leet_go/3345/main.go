package main

import (
	"fmt"
)

func main() {
	fmt.Println("Result: ", smallestNumber(10, 2))
}

func smallestNumber(n int, t int) int {
	result := n
	for {
		result = digitProduct(n)
		fmt.Println("result: ", result)
		if result%t == 0 {
			break
		}

		n++
	}

	return n
}

func digitProduct(num int) int {
	product := 1
	for num > 1 {
		fmt.Println("num: ", num)
		digit := num % 10
		product *= digit
		num /= 10
	}

	return product
}
