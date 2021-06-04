package main

import (
	"fmt"
)

func foo() func(x int) int {
	return func(x int) int {
		return 770 + x
	}
}

func main() {
	fun := foo()

	fmt.Print(fun(30))
}
