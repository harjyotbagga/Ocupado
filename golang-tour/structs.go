package main

import (
	"fmt"
	"strconv"
)

// Defining Struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int
	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello! " + p.firstName
}

func (p Person) askAge() string {
	return "Hey! My age is " + strconv.Itoa(p.age)
}

// hasBirthday (pointer revceiver)
func (p *Person) hasBirthday() {
	p.age++
}

func main() {
	person1 := Person{
		firstName: "Harjyot",
		lastName:  "Bagga",
		city:      "Vellore",
		gender:    "M",
		age:       20}

	fmt.Println(person1)
	fmt.Println(person1.firstName)
	fmt.Println()
	fmt.Println(person1.greet())
	fmt.Println(person1.askAge())
	person1.hasBirthday()
	fmt.Println(person1)

}
