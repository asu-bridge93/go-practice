package math

import (
	"fmt"
	"math/rand"
)

func RandomNumber() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
