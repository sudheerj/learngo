package main

import "fmt"

//Variable declaration
var msg string

func main() {
	//Assignment
	msg = "Good Morning"
	fmt.Println(msg)

	//Default value without initialization
	var zero int
	fmt.Println(zero)

	//Multiple variables declaration
	var one, two int = 1, 2
	fmt.Println(one, two)

	//Shorthand notation for declaration and initialization
	three:= 3
	fmt.Println(three)
}
