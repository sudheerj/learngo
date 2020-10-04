package main

import (
	"fmt"
)

type User struct {
	name     string
	age   int
	city string
	country string
}


//Use User pointer as the receiver type
func (u *User) changeCity(city string) {
	u.city = city
}

// displayAddress() method uses User as the receiver type
func (u User) displayAddress() {
	fmt.Printf("Address of %s is %s, %s \n", u.name, u.city, u.country)
}

func main() {
	user := User {
		name:     "John",
		age:   30,
		city: "NewYork",
		country: "US",
	}
	user.displayAddress()
	user.changeCity("NewJersey") //(&user).changeCity("NewJersey")
	user.displayAddress()
}