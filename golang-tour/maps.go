package main

import (
	"fmt"
)

func main() {
	emails := make(map[string]string)
	emails["Bob"] = "bob@gmail.com"
	emails["Jenny"] = "jenny@gmail.com"
	emails["Mike"] = "mike@gmail.com"
	fmt.Println(emails)
	fmt.Println(emails["Bob"])
	fmt.Println(len(emails))

	delete(emails, "Jenny")
	fmt.Println(emails)

	age := map[string]int{"A": 20, "B": 18}
	fmt.Println(age)
}
