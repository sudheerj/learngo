package main

import "fmt"

func main() {
	var i interface{} = "Welcome"

	// type switch is used to find the interface{} type
	switch t := i.(type){
	case int64:
		fmt.Println("Integer type:", t)
	case float64:
		fmt.Println("Float type:", t)
	case string:
		fmt.Println("String type:", t)
	case bool:
		fmt.Println("Bool type:", t)
	case nil:
		fmt.Println("nil type")
	default:
		fmt.Println("Unknown type!")
	}
}
