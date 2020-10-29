package main

import (
	"fmt"
)

func goRoutineOne(ch chan int) {
	ch <- 100
}

func goRoutineTwo(ch chan int) {
	ch <- 200
}

func goRoutineThree(ch chan int) {
	ch <- 300
}

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go goRoutineOne(ch1)
	go goRoutineTwo(ch2)
	go goRoutineThree(ch3)

	select {
	case value1 := <-ch1:
		fmt.Println("First go routine printed: ", value1)
	case value2 := <-ch2:
		fmt.Println("Second go routine printed:", value2)
	case value3 := <-ch3:
		fmt.Println("Third go routine printed:", value3)
	default:
		fmt.Println("The default case!!")
	}
}