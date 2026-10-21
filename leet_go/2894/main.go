package main

import "fmt"

func main() {
	differenceOfSums(10, 3)
}

func differenceOfSums(n int, m int) int {

	result := 0
	for i := 0; i < n; i++ {
		currNum := i + 1
		if currNum%3 != 0 {
			fmt.Println("add: ", currNum)
			result += currNum
		} else {
			fmt.Println("remove: ", currNum)
			result -= currNum
		}
	}

	fmt.Print("result: ", result)

	return 0
}
