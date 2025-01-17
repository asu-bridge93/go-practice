package functions

import (
	"fmt"
)

func Add(x int, y int) int {
	fmt.Printf("%d + %d = %d.\n", x, y, x+y)
	return x + y
}
