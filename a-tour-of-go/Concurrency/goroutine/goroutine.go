package goroutine

import (
	"fmt"
	"time"
)

func GoRoutine(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
