package main

import "fmt"

func main() {
	fmt.Print(
		"Result: ", findRestaurant(
			[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
			[]string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"},
		),
	)
}

func findRestaurant(list1 []string, list2 []string) []string {
	mapList2 := make(map[string]int)
	for i, value := range list2 {
		mapList2[value] = i
	}

	minIndex := len(list1) + len(list2)

	var result []string

	for i, value := range list1 {
		if indexList2, exist := mapList2[value]; exist {
			tempIndex := i + indexList2
			if tempIndex < minIndex {
				minIndex = tempIndex
				result = []string{value}
			} else if tempIndex == minIndex {
				result = append(result, value)
			}
		}
	}
	return result
}

func findRestaurant1(list1 []string, list2 []string) []string {

	sameString := make(map[int][]string)
	leastIndex := len(list1) + len(list2)
	for i, value1 := range list1 {
		for j, value2 := range list2 {
			if value1 == value2 {
				sumIndex := i + j
				if leastIndex > sumIndex {
					leastIndex = sumIndex
				}
				fmt.Println("SAME: ", value1, " at i:", i, ", j:", j, ", total index: ", sumIndex)
				value, _ := sameString[sumIndex]
				sameString[sumIndex] = append(value, value1)
			}
		}
	}

	if len(sameString) == 0 {
		return []string{}
	}

	fmt.Println("MAP: ", sameString, ", leastIndex: ", leastIndex)

	return sameString[leastIndex]
}
