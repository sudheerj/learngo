package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3)

	//producer
	go producer(c)

	//consumer
	go consumer(c)
}

func producer(ch chan int) {
	fmt.Println("producer started")
	for i := 1; i <= 3; i++ {
		ch <- i
	}
	close(ch)
	fmt.Println("producer finished")
}

func consumer(ch chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("consumer started")
	for i := range ch {
		fmt.Println("i =", i)
	}
	fmt.Println("consumer finished")
}
