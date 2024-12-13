package main

import "fmt"

func eachCons(list []int, n int) [][]int {
	var result [][]int
	if n <= 0 || n > len(list) {
		return result
	}

	for i := 0; i <= len(list)-n; i++ {
		subset := list[i : i+n]
		result = append(result, subset)
	}

	return result
}

func main() {
	fmt.Println(eachCons([]int{1, 2, 3, 4, 5, 6}, 2))
}
