package channel

func Channel(s []int, channel chan int) {
	sum := 0
	for v := range s {
		sum += v
	}
	channel <- sum // send sum to channel
}
