package main

import "fmt"

// Create a struct Rectangle with fields Length and Width (both float64).
//  Write a method to calculate and return the area of the rectangle.

type Rectangle struct {
	Length float64
	Width  float64
}

func getArea(length float64, width float64) string {
	rectangle := Rectangle{
		Length: length,
		Width:  width,
	}
	area := rectangle.Length * rectangle.Width
	sentence := fmt.Sprintf("The area of a rectangle of length %v and breadth %v is %v",
		rectangle.Length, rectangle.Width, area)
	return sentence
}

func main() {
	rectangle1 := getArea(13.2, 12)
	rectangle2 := getArea(13, 12.2)
	rectangle3 := getArea(13.2, 12.2)
	fmt.Println(rectangle1, rectangle2, rectangle3)
}
