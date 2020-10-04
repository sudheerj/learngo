package main

import (
	"fmt"
)

type UserProfile struct {
	name     string
	age   int
	city string
	country string
}


// displayAddress() method uses User as the receiver type
func (u UserProfile) displayAddress() {
	fmt.Printf("Address of %s is %s, %s \n", u.name, u.city, u.country)
}

func main() {
	user := User {
		name:     "John",
		age:   30,
		city: "NewYork",
		country: "US",
	}
	user.displayAddress() //Call displayAddress() method of User type
}