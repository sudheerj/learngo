package main

import "fmt"

func add(a, b int) {
	sum := a + b
	fmt.Println("Sum: ", sum)
}

func message(msg string) {
	fmt.Println(msg)
}

func main() {

	message("Begin")

	defer message("End")
	defer add(10, 20)
	defer add(20, 30)
	defer add(30, 40)
}
