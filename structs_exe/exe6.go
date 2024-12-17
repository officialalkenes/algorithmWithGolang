package main

// Define a struct `Employee` with fields `Name`, `Position`, and `Salary`.
// Write a function to give an employee a raise by a percentage of their current salary.
import (
	"fmt"
)

type Employee struct {
	Name     string
	Position string
	Salary   float32
}

func giveEmployeeRaise(employee Employee) float32 {
	salary := employee.Salary
	percentage := salary * 0.1
	final_balance := salary + percentage
	return final_balance
}

func main() {
	employees := []Employee{
		{Name: "Olamide", Position: "CTO", Salary: 9000000},
		{Name: "Desmond", Position: "Manager", Salary: 900000},
		{Name: "Jazzy", Position: "Lead Agent", Salary: 900000},
		{Name: "Jamiu", Position: "Cheif Security", Salary: 100000},
	}
	for _, employee := range employees {
		new_wage := giveEmployeeRaise(employee)
		employee.Salary = new_wage
		fmt.Printf("The new salary of %s. who was promoted to %s is now %v\n", employee.Name, employee.Position, employee.Salary)
	}
}
