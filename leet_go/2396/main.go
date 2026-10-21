package main

import (
	"fmt"
	util "leet_go/util"
	"strconv"
)

func main() {
	//fmt.Println(isStrictlyPalindromic(9))
	wew := []int{1, 2, 3, 3, 5, 5}
	wew = append(wew, 7)
	fmt.Println(util.RemoveAll(wew, 3))
	fmt.Println(util.Clear(wew))
	//fmt.Println(isPalindrome("abcba"))
}

func isStrictlyPalindromic(n int) bool {
	for i := 2; i < n-1; i++ {
		if !isPalindrome(
			strconv.FormatInt(int64(n), i),
		) {
			return false
		}
	}

	return true
}

func isPalindrome(word string) bool {
	var j int = len(word) - 1
	for i := 0; i < len(word)-1; i++ {
		iChar := word[i]
		jChar := word[j]
		if iChar != jChar {
			return false
		}
		j--
	}

	return true
}
