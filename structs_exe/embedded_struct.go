package main

import "fmt"

// Define the Car struct
type Car struct {
	Make  string
	Model string
}

// Define the Truck struct that embeds Car
type Truck struct {
	Car
	BedSize int
	Wheels  int
}

func main() {
	// Create an instance of Truck
	truck1 := Truck{
		Car: Car{ // Initialize the embedded Car struct
			Make:  "Fiat",
			Model: "Model 0",
		},
		BedSize: 13,
		Wheels:  19,
	}

	// Access fields of Truck and Car
	fmt.Println("Truck Details:")
	fmt.Printf("Make: %s\n", truck1.Make)   // Accessed via embedded Car
	fmt.Printf("Model: %s\n", truck1.Model) // Accessed via embedded Car
	fmt.Printf("Bed Size: %d\n", truck1.BedSize)
	fmt.Printf("Wheels: %d\n", truck1.Wheels)
}
