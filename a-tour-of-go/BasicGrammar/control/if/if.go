package controlif

import (
	"fmt"
	"math"
)

func IfStatement(x float64) string {
	if x < 0 {
		result := IfStatement(-x) + "i"
		fmt.Println(result)
		return result
	}
	return fmt.Sprint(math.Sqrt(x))
}
