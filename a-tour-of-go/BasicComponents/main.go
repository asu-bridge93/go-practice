package main

import (
	"fmt"
	"go-practice/constants"
	"go-practice/functions"
	"go-practice/math"
	"go-practice/prints"
	"go-practice/types"
	"go-practice/variables"
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

	variables.Bool()
	variables.InitVariables()
	variables.ShortInitVariables()

	types.BasicTypes()
	variables.ZeroValues()
	types.TypeConversion()
	types.TypeInference()

	constants.Constants()
	constants.NumValues()
}
