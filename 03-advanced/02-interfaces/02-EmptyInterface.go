package main

import "fmt"

func main() {
	var i interface{}
	type User struct {
		name string
	}
	print(i)

	i = 10
	print(i)

	i = "Welcome"
	print(i)

	i = true
	print(i)

	i = User{"John"}
	print(i)
}

func print(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

