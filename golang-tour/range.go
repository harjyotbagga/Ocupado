package main

import (
	"fmt"
)

func main() {
	ids := []int{5, 10, 15, 20, 25, 30}
	for i, id := range ids {
		fmt.Printf("%d - ID No. %d\n", i, id)
	}

	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Printf("Sum is %d\n", sum)

	emails := make(map[string]string)
	emails["Bob"] = "bob@gmail.com"
	emails["Jenny"] = "jenny@gmail.com"
	emails["Mike"] = "mike@gmail.com"
	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

}
