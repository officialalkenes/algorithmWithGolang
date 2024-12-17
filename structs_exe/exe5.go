package main

import (
	"fmt"
)

// Implement a struct `Car` with fields `Make`, `Model`, and `Year`.
// Write a function to check if a car is a classic (more than 25 years old).

type Car struct {
	Make  string
	Model string
	Year  int16
}

func checkIsClassic(car Car) bool {
	if 2024-car.Year > 25 {
		return true
	} else {
		return false
	}
}

func main() {
	cars := []Car{
		{Make: "Toyota", Model: "Camry 09", Year: 2009},
		{Make: "Mercedes", Model: "Benz", Year: 2009},
		{Make: "Volkswagen", Model: "Turtle", Year: 1980},
	}
	for _, car := range cars {
		get_value := checkIsClassic(car)
		if get_value {
			fmt.Printf("%s is a Classic car. It was made %v years ago\n", car.Model, 2024-car.Year)
		} else {
			fmt.Printf("%s is not a Classic car. It was made %v years ago\n", car.Model, 2024-car.Year)
		}
	}
}
