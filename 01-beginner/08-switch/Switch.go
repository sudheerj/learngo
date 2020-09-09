package main

import (
	"fmt"
	"time"
)

func main() {
	switch hour := time.Now().Hour();{
	case hour < 12 :
		fmt.Println("Good morning!!")
	case hour < 15:
		fmt.Println("Good Afternoon!!")
	default:
		fmt.Println("Good evening!!")
	}
}
