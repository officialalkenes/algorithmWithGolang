package main

import "fmt"

// Define a struct Person with fields Name (string),
//  Age (int), and PhoneNumber (string).
// Write a function to print a formatted description of a Person.

type Person struct {
	Name        string
	Age         int
	PhoneNumber string
}

func main() {
	person1 := Person{
		Name:        "Olamide",
		Age:         12,
		PhoneNumber: "08012345678",
	}

	sentence := fmt.Sprintf("My name is %s. I am %d years old. Here is my number: %s", person1.Name, person1.Age, person1.PhoneNumber)
	fmt.Println(sentence)
}
