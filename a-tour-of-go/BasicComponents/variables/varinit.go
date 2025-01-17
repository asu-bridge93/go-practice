package variables

import "fmt"

var i, j int = 1, 2

func InitVariables() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
