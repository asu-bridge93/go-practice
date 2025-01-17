package control

import "fmt"

func SumForloop(n int) {
	sum := 0
	for i := 0; i < n; i++ {
		sum += i
	}
	fmt.Println(sum)
}
