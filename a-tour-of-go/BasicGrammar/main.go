package main

import (
	"fmt"
	"go-practice/control"
	controldefer "go-practice/control/defer"
	controlfor "go-practice/control/for"
	controlif "go-practice/control/if"
	controlswitch "go-practice/control/switch"
	controlwhile "go-practice/control/while"
)

func main() {
	controlfor.SumForloop(15)
	controlfor.ShortSumForloop(1024)
	controlwhile.ForisGosWhile(10000)
	controlif.IfStatement(400)
	controlif.IfStatement(-100)
	fmt.Println(
		controlif.Pow(5, 3, 150),
		controlif.Pow(2, 4, 20),
	)
	fmt.Println(
		controlif.IfandElse(3, 2, 10),
		controlif.IfandElse(3, 3, 20),
	)
	fmt.Println(
		control.NewtonSqrt(1),
		control.NewtonSqrt(2),
		control.NewtonSqrt(3),
		control.NewtonSqrt(4),
	)

	controlswitch.Switch()
	controlswitch.SwitchDay()
	controlswitch.Greeting()

	controldefer.Defer()
	controldefer.DeferMulti()
}
