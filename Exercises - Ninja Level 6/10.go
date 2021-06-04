package main

import (
	"fmt"
)

func foo() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
func main() {
	fun := foo()
	fmt.Println(fun())
	fmt.Println(fun())
	fmt.Println(fun())
	fmt.Println(fun())
}
