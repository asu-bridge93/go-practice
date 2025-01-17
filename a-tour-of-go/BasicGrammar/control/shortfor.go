package control

import "fmt"

func ShortSumForloop(n int) {
	sum := 1      // ここで初期化しているので、
	for sum < n { // 宣言しなくていい
		sum += sum
	}
	fmt.Println(sum)
}
