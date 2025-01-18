package controlwhile

import "fmt"

func ForisGosWhile(n int) {
	sum := 1
	for sum < n {
		sum += sum
	}
	fmt.Println(sum)
}
