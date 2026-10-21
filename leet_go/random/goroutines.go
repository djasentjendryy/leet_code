package main

import (
	"fmt"
)

func sayHi() {
	fmt.Println("hi")
}

func main() {
	go fmt.Println("you'll never see me")
}
