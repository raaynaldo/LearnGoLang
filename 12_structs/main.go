package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	// firstname string
	// lastname  string
	// city      string
	// gender    string
	// age       int

	firstname, lastname, city, gender string
	age                               int
}

// Greeting method(value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstname + " " + p.lastname + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(SpouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastname = SpouseLastName
	}
}

func main() {
	// Init person using struct
	person1 := Person{firstname: "Felicia", lastname: "Hanslim", city: "Jakarta", gender: "f", age: 22}
	// Init person using struct
	person2 := Person{"Teofilus", "Raynaldo", "Jakarta", "m", 24}

	// fmt.Println(person1.firstname)
	// person1.age++
	// fmt.Println(person1)

	person1.hasBirthday()
	person1.getMarried("Raynaldo")

	person2.getMarried("Cynthia")

	fmt.Println(person1.greet())
	fmt.Println(person2.greet())

}
