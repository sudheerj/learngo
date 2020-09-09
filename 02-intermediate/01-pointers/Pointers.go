package main

import "fmt"

func main() {
	//Variable declaration
	a := 10

	//Pointer reference to variable 'a'
	p := &a
	var p1 *int = &a

	//Find the memory address of variable 'a'
	fmt.Println(&a)
	fmt.Println(p)
	fmt.Println(p1)

	//Read variable value
	fmt.Println(*p)

	//Write variable value
	*p = 20
	fmt.Println(a)
	fmt.Println(*p)
}
