package main

import "fmt"

// 4. Create a struct `Student` with fields `Name`,
// `Grades` (slice of ints), and `AverageGrade`.
// Write a function to calculate and update the average grade for a `Student`.

type Student struct {
	Name         string
	Grades       []int
	AverageGrade float64
}

func calculateAvg(student Student) float64 {
	grades := student.Grades
	total := 0
	for _, grade := range grades {
		total = total + grade
	}
	average := float64(total) / float64(len(grades))
	return float64(average)
}

func main() {
	student1 := Student{
		Name:   "Olamide",
		Grades: []int{32, 44, 92},
	}
	avg_grade := calculateAvg(student1)
	student1.AverageGrade = avg_grade
	fmt.Printf("The name of the new student is %s. Here are his grades: %v. His Average grade is %.2f.\n",
		student1.Name, student1.Grades, student1.AverageGrade)
}
