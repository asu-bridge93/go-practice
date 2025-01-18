package channel

func Close(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x // channelに値を送信
		x, y = y, x+y
	}
	close(c) // これ以上送信する値ことがないことを示すため、channelをclose
}
