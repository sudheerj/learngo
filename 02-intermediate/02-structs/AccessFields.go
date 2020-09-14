package main

import "fmt"

type AddressInfo struct {
	name, street, city, country string
	pincode int
}

func main() {
	// Declaring a struct literal
	address := AddressInfo{"John", "CommonWealth", "Singapore", "Singapore", 140098}
	fmt.Println("Name: ", address.name)
	fmt.Println("City: ", address.street)

	address.street = "Queenstown"
	fmt.Println("New address: ", address)
}