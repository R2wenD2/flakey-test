package main

import (
	"fmt"

	"github.com/R2wenD2/flakey-test/calc"
)

func main() {
	num := calc.Add(1, 1)
	fmt.Sprintf("The result is %d", num)
}
