package main

import "fmt"

func main() {
	even := []int{0, 2, 4, 6, 8, 10, 12}

	for index, value := range even {
		fmt.Printf("Even[%d] = %d\n", index, value)
	}
}

