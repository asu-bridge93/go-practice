package controlif

import (
	"fmt"
	"math"
)

func IfandElse(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
		return lim
	}
}
