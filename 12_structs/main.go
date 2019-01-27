package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	// firstName string
	// lastName  string
	// age       int
	// city      string
	// gender    string

	firstName, lastName, city, gender string
	age                               int
}

// Greeting method(value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method(pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

func main() {
	// Initials Person With Struct
	person1 := Person{firstName: "Mohammad", lastName: "Katoozi", age: 21, city: "Tehran", gender: "f"}

	// Alternative
	// person2 := Person{"Mohammad", "Katoozi", 21, "Tehran", "m"}

	// fmt.Println(person1, person2)
	// fmt.Println(person1.firstName, person2.firstName)
	// person2.firstName = "Ali"
	// fmt.Println(person1.firstName, person2.firstName)
	fmt.Println(person1.greet())
	person1.hasBirthday()
	fmt.Println(person1.greet())
}
