package main

import (
	"fmt"
	"math"
)

func shortestCompletingWord(licensePlate string, words []string) string {
	var requireAlphabet [26]int
	for _, v := range licensePlate {
		if !(v >= 65 && v <= 90) && !(v >= 97 && v <= 122) {
			continue
		}

		if v <= 90 {
			requireAlphabet[v-65]++
		} else {
			requireAlphabet[v-97]++
		}
	}

	shortestCWord := ""
	for _, word := range words {
		var alphabetIndex [26]int

		for _, w := range word {
			alphabetIndex[w-'a']++

		}

		doesMatched := true
		for i := 0; i < 26; i++ {
			if requireAlphabet[i] > alphabetIndex[i] {
				doesMatched = false
				break
			}
		}

		if doesMatched {
			if len(word) < len(shortestCWord) || shortestCWord == "" {
				shortestCWord = word
			}
		}

	}
	return shortestCWord
}

func shortestCompletingWord1(licensePlate string, words []string) string {
	var requireAlphabet [26]int
	for _, v := range licensePlate {

		if !(v >= 65 && v <= 90) && !(v >= 97 && v <= 122) {
			continue
		}

		if v <= 90 {
			requireAlphabet[v-65]++
		} else {
			requireAlphabet[v-97]++
		}
	}

	result := words[0]
	tempLen := math.MaxInt
	tempUnMatchedAlphabet := math.MaxInt

	for _, word := range words {
		tempRequiredALphabet := requireAlphabet
		for _, w := range word {
			if w >= 65 && w <= 90 {
				tempRequiredALphabet[w-65]--
			} else {
				tempRequiredALphabet[w-97]--
			}
		}

		unMatchedAlphabet := 0
		for _, vALphabet := range tempRequiredALphabet {
			if vALphabet > 0 {
				unMatchedAlphabet += vALphabet
			}
		}

		fmt.Println("Word: ", word, "unMatchedAlphabet: ", unMatchedAlphabet, "array Matched: ", tempRequiredALphabet)
		if tempUnMatchedAlphabet > unMatchedAlphabet {
			tempUnMatchedAlphabet = unMatchedAlphabet
			result = word
			tempLen = len(word)

		} else if tempUnMatchedAlphabet == unMatchedAlphabet && len(word) < tempLen {
			tempUnMatchedAlphabet = unMatchedAlphabet
			result = word
			tempLen = len(word)

		}

	}

	return result
}

func main() {
	fmt.Println("result: ", shortestCompletingWord("1s3 PSt", []string{"step", "steps", "stripe", "stepple"}))
}
