package main

// Create a struct `Point` with fields `X` and `Y` (both int).
// Write a function to calculate the distance between two points.

import (
	"fmt"
)

type Point struct {
	X float32
	Y float32
}

func getDistant(point Point) float32 {
	distance := point.X - point.Y
	return distance
}

func main() {
	coordinates := []Point{
		{X: 900, Y: 700},
		{X: 900, Y: 192.34},
		{X: 800.4, Y: 700},
	}
	for _, cordinate := range coordinates {
		distance := getDistant(cordinate)
		fmt.Printf("The distance between %v and %v is %v\n", cordinate.X, cordinate.Y, distance)
	}
}
