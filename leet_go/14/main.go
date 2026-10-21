package main

import (
	"fmt"
)

func main() {
	longestCommonPrefix([]string{"flower", "flow", "flight"})
}

func longestCommonPrefix(strs []string) string {
	currentPrefix := strs[0]
	lastIndex := len(strs[0])
	strs = strs[1:]
	fmt.Println("Strs: ", strs)
	for i, v := range strs {
		for j := 0; j < len(v) && j < len(currentPrefix); j++ {
			if currentPrefix[j] != v[j] {
				fmt.Println("beda niiih -> ", j, currentPrefix[j])
				lastIndex = j
				break
			}
		}
		currentPrefix = currentPrefix[:lastIndex]
		fmt.Println(i, v, ", last index: ", lastIndex, ", current prefix: "+currentPrefix)
	}
	return ""
}
