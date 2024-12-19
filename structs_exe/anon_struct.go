package main

import "fmt"

func main() {
	myCar := struct {
		Make  string
		Model string
	}{
		Make:  "Tesla",
		Model: "Model 3",
	}
	fmt.Println(myCar)
}
