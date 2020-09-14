package main

import "fmt"

type User struct {
	firstName, lastName string
	age, pincode int
}

func main() {
	user := &User{"John", "Anderson", 30, 140098}
	fmt.Println((*user).firstName)
	fmt.Println(user.firstName)
	user.firstName = "Jack"
	fmt.Println(user.firstName)
}
