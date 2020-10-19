package main

import "fmt"

func main() {
	// Arrays
	var fruitArr [2]string
	fruitArr[0] = "Apple"
	fruitArr[1] = "Banana"

	fruitArr1 := [2]string{"Mango", "Orange"}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])
	fmt.Println(fruitArr1)

	fruitSlice := []string{"Apple", "Mango", "Orange", "Grapes"}
	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])

}
