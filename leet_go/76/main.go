package main

func main() {
	minWindow("ADOBECODEBANC", "ABC")
}

func minWindow(s string, t string) string {
	var requiredElementUpper [26]int
	for i := 0; i < len(t); i++ {
		requiredElementUpper[t[i]-'A']++
	}
	// var countElementUpper [26]int

	for i := 0; i < len(s); i++ {

	}

	return ""
}
