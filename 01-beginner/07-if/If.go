package main

import "fmt"

func oddOrEven(x int) {
	if v := x%2; v == 0 {
		fmt.Printf("The reminder is %v. Number %v is an even number\n", v,x)
	} else {
		fmt.Printf("The reminder is %v. Number %v is an odd number\n", v,x)
	}
}

func main() {
		oddOrEven(10)
		oddOrEven(5)
}
