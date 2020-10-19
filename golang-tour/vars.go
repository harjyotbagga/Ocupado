package main

import (
	"fmt"
)

func main() {
	var name string = "HSB"
	var age = 20
	var isTrue = true
	new_name := "Golang"

	user_name, email := "Harjyot", "hsbagga@xyz.com"

	fmt.Println(name, age)
	fmt.Printf("%T\n", isTrue)
	fmt.Println(new_name)
	fmt.Printf("%T\n", new_name)
	fmt.Println(user_name, email)

}
