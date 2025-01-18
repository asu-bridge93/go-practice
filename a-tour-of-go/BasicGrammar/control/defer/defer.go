package controldefer

import "fmt"

func Defer() {
	defer fmt.Println("World")

	fmt.Println("Hello") // このreturn文が呼び出されるまで実行されない。
}
