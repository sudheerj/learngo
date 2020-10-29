package main

import (
	"fmt"
	"time"
)

func myFunc(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	myFunc("Regular function call")

	go myFunc("Goroutine call")

	go func(msg string) {
		fmt.Println(msg)
	}("Anonymous call")

	time.Sleep(5 * time.Second)

	fmt.Println("Completed")
}