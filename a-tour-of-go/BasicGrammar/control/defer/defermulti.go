package controldefer

import "fmt"

func DeferMulti() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Printf("%d\n", i)
	}

	fmt.Println("done")
}
