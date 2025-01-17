package constants

import "fmt"

const Pi = 3.14

func Constants() {
	const world = "世界" // := を使っての定義は出来ない。
	fmt.Println("Hello", world)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
