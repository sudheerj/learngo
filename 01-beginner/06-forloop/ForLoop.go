package main

import "fmt"

func main() {
	//Three components(init statement, condition expression and post statement)
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	//Optional init and post statements(Works as while loop in other languages)
	j := 1
	for j <= 5 {
		fmt.Println(j)
		j = j + 1
	}

	//No conditions. Becomes infinite loop without break statement
	for {
		fmt.Println("Break infinite loop")
		break
	}

	//Skip iterations with continue keyword. e.g, Skip odd numbers
	for k := 1; k <= 5; k++ {
		if k%2 != 0 {
			continue
		}
		fmt.Println(k)
	}
}