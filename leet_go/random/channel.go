package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int, 1000)

	for i := 0; i < 1000; i++ {
		c <- i
	}

	go func() {
		for msg := range c {
			fmt.Println("message is %d", msg)
		}
	}()

	time.Sleep(1 * time.Second)

	// ch := make(chan int) // a channel of int
	// ch <- 42             // send (the arrow points INTO the channel)
	// val := <-ch          // receive (the arrow points OUT)
	// fmt.Println(val)
	// close(ch) // close it (sender's responsibility)

	// c := make(chan int)
	// go func() {
	// 	for msg := range c {
	// 		fmt.Println("message is ", msg)
	// 	}
	// }()

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("i: ", i)
	// 	c <- i
	// 	time.Sleep(5 * time.Second)
	// }
	// fmt.Println("BYE")

	// time.Sleep(2 * time.Second)
}

// func main() {
// 	c := make(chan int, 1)

// 	// go func() {
// 	// 	for {
// 	// 		fmt.Println("start sleep")
// 	// 		time.Sleep(1 * time.Second)
// 	// 		fmt.Println("stop sleep")
// 	// 	}
// 	// }()
// 	fmt.Println("======")
// 	for msg := range c {
// 		fmt.Println("message is %d", msg)
// 	}
// 	fmt.Println("======1")

// 	go func() {
// 		for i := 0; i < 1; i++ {
// 			c <- i
// 		}
// 	}()

// 	time.Sleep(1 * time.Second)
// 	fmt.Print("======")
// }
