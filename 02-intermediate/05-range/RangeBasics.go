package main

import "fmt"

func main() {
	even := [7]int{0, 2, 4, 6, 8, 10, 12}
	odd := []int{1, 3, 5, 7, 9, 11, 13}
	msg := "HelloWorld"
	numbers := map[string]int{"One": 1,
		"Two": 2, "Three": 3}


	for index, value := range even {
		fmt.Printf("Even[%d] = %d\n", index, value)
	}

	for index, value := range odd {
		fmt.Printf("Odd[%d] = %d\n", index, value)
	}

	for index, value := range msg {
		fmt.Printf("Msg[%d] = %d\n", index, value)
	}

	for key, _ := range numbers {
		fmt.Println("Number of", key, "is: ", numbers[key])
	}
}

