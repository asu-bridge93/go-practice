package channel

import "fmt"

func Buffer() {
	ch := make(chan int, 2) // Buffer付きchannel
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
