package main

import "fmt"

func main() {
	lengthOfLongestSubstring("abcabcbb")
}

func lengthOfLongestSubstring(s string) int {
	fmt.Println(s)
	var result string

	i := 1
	j := 0

	result += string(s[j])

	for i <= len(s) {
		fmt.Println(result)

		i++
	}

	result += "a"
	result = ""
	fmt.Println("apkdjalksdjlakjslkdjaskd ====")
	fmt.Println(result)
	return 0
}
