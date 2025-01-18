package channel

import "fmt"

func Select(ch, quit chan int) {
	x, y := 0, 1 // Fibonacci 数列の初期値
	for {
		select {
		case ch <- x:
			// ch チャネルに x を送信
			// 次の Fibonacci 数値を計算
			x, y = y, x+y
		case <-quit:
			// quit チャネルから受信があった場合
			fmt.Println("quit")
			return
		}
	}
}
