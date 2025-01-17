package main

import (
	"fmt"
	"go-practice/functions"
	"go-practice/math"
	"go-practice/prints"
)

func main() {
	prints.PrintHelloWorld()
	prints.Showtime()

	math.RandomNumber()
	math.SqrtNumber()
	math.Pi()

	functions.Add(3, 5)
	functions.Subtract(5, 2)
	school, department := functions.Swap("The University of Tokyo", "PSI")
	fmt.Println(school, department)
	fmt.Println(functions.NakedReturn(17))
}
