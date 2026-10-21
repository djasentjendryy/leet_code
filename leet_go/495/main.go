package main

import "fmt"

func main() {
	fmt.Print("result : ", findPoisonedDuration([]int{1, 4}, 2))
}

/**
second -> 1, 2 ; duration 3
1  ->
.2 ->
3
**/

func findPoisonedDuration(timeSeries []int, duration int) int {
	if len(timeSeries) == 0 {
		return 0
	}

	countTime := duration
	for i := 1; i < len(timeSeries); i++ {
		gap := timeSeries[i] - timeSeries[i-1]

		if gap > duration {
			countTime += duration
		} else {
			countTime += gap
		}
	}
	return countTime
}

func findPoisonedDuration1(timeSeries []int, duration int) int {
	if len(timeSeries) == 0 {
		return 0
	}
	totalTime := make(map[int]struct{})
	countTime := 0
	for _, value := range timeSeries {
		for i := range duration {
			second := value + i
			if _, ok := totalTime[second]; !ok {
				countTime++
			}
			totalTime[second] = struct{}{}
		}
	}
	return countTime
}
