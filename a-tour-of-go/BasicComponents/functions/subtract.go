package functions

import "fmt"

func Subtract(x, y int) int {
	fmt.Printf("%d - %d = %d.\n", x, y, x-y)
	return x - y
}
