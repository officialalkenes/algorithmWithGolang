package main

import "fmt"

type RectangleArea struct {
	Length  float64
	breadth float64
}

func (r RectangleArea) getRecArea() float64 {
	return r.breadth * r.Length
}

func main() {
	rect1 := RectangleArea{
		Length:  101.4,
		breadth: 12.3,
	}
	fmt.Println(rect1.getRecArea())
}
