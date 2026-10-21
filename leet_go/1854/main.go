package main

import "fmt"

func maximumPopulation(logs [][]int) int {

	diff := make([]int, 101)
	for _, v := range logs {
		diff[v[0]-1950]++
		diff[v[1]-1950]--
	}

	fmt.Println(diff)
	maxPop, year, currPop := 0, 1950, 0
	for i, v := range diff {
		currPop += v
		if maxPop < currPop {
			maxPop = currPop
			year = 1950 + i
		}
	}

	return year
}

func main() {
	fmt.Println("Result: ", maximumPopulation([][]int{{1950, 1961}, {1960, 1971}, {1970, 1981}}))
}
