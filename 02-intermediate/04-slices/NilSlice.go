package main

import "fmt"

func main() {
	var nilSlice []int
	fmt.Println(nilSlice, len(nilSlice), cap(nilSlice))
}
