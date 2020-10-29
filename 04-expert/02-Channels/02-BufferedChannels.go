package main

import "fmt"

func main() {
	ch := make(chan string, 2)
	ch <- "GoodMorning!"
	ch <- "John"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
