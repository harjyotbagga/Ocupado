package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b)
}
