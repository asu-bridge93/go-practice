package main

import (
	"fmt"
	"goroutine/channel"
	"goroutine/goroutine"
)

func main() {
	// goroutine
	go goroutine.GoRoutine("Hello")
	goroutine.GoRoutine("World") // 同時に複数のスレッドで実行される

	// channel
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int) // channelは一つでも大丈夫
	go channel.Channel(s[:len(s)/2], c)
	go channel.Channel(s[len(s)/2:], c) // 二つのgoroutineの間で作業を分配
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)

	channel.Buffer()

	c = make(chan int, 30)
	go channel.Close(cap(c), c)
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\n")

	// Close関数の影響で動かないため、コメントアウト
	// c = make(chan int, 20)
	// quit := make(chan int)
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(<-c)
	// }
	// quit <- 0
	// channel.Select(c, quit)

}
