package variables

import "fmt"

func ShortInitVariables() {
	var i, j int = 1, 2
	k := 3 // var宣言の代わり。暗黙的な型宣言ができる
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
