package main

import "fmt"

func main() {
	fmt.Println("Result: ", convertToTitle(27))
}

func convertToTitle(columnNumber int) string {

	n := columnNumber
	var result string

	for n > 0 {
		n--
		remain := n % 26
		result = string(rune('A'+remain)) + result
		n = n / 26
	}
	return result
}

/**

A -> 1
B -> 2
C -> 3
...
Z -> 26
AA -> 27
AB -> 28
...
BYFN -> 52222
**/
