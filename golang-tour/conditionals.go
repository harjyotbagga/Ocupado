package main

import "fmt"

func main() {

	x := 10
	y := 15
	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else if y > x {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is equal to %d\n", x, y)
	}

	color := "green"
	switch color {
	case "red":
		fmt.Println("Color is red.")
	case "blue":
		fmt.Println("Color is blue")
	default:
		fmt.Println("Color is neither")
	}

}
