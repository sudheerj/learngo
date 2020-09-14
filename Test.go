package main

import "fmt"

var gvar int //not an error

func main() {
	var one int   //error, unused variable
	two := 2      //error, unused variable
	var three int //error, even though it's assigned 3 on the next line
	three = 3

	func(unused string) {
		fmt.Println("Unused arg. No compile error")
	}("what?")
}