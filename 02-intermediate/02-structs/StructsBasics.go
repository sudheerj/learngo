package main

import "fmt"

type Address struct {
	name, street, city, country string
	pincode int
}

func main() {
		// By default struct fields are initialized with their zero value
		var addressOne Address
		fmt.Println("Empty address:", addressOne)

		// Declaring a struct literal
		addressTwo := Address{"John", "CommonWealth", "Singapore", "Singapore", 140098}
		fmt.Println("Second address: ", addressTwo)

		// Naming fields of a struct
		addressThree := Address{name: "Rock", country: "England", pincode: 100084}
		fmt.Println("Third address:", addressThree)

		// Uninitialized fields are set to zero-value
		addressFour := Address{name: "Warner", country: "Australia"}
		fmt.Println("Fourth address: ", addressFour)
}