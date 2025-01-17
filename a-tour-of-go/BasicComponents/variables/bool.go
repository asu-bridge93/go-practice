package variables

import "fmt"

var c, python, java bool // 初期値として false, false, false

func Bool() {
	var i int // 初期値として 0
	fmt.Println(i, c, python, java)
}
