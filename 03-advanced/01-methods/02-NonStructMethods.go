package main

import "fmt"

type myInt int

func (x myInt) add(y myInt) myInt {
	return x + y
}

func main() {
	num1 := myInt(10)
	num2 := myInt(15)
	sum := num1.add(num2)
	fmt.Printf("Sum of %d and %d is %d \n", num1, num2, sum)
}