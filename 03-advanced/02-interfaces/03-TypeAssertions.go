package main

import "fmt"

func main() {
	var i interface{} = 100

	//Retrieves the int value from interface
	x := i.(int)
	fmt.Println(x)

	//Tests if interface has string type and return true if found otherwise false
	y, ok := i.(string)
	if ok {
		fmt.Printf("String value exists! The value is %v", y)
	} else {
		fmt.Println("No string value available!! Let's print zero value of the string i.e, empty string")
	}

	//Return panic as interface does not have float64 type
	z := i.(float64)
	fmt.Println(z)
}
