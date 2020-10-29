package main

import "fmt"

func main() {
	ch := make(chan string, 2)
	ch <- "GoodMorning!"
	ch <- "John"
	ch <- "How are you!!"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
