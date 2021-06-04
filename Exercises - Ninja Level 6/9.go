package main

import (
	"fmt"
)

func foo(x []int) int {
	return x[0]
}
func bar(foo func(a []int) int, y int) {
	fmt.Println(foo([]int{1, 2, 3, 4}) * y)
}

func main() {
	bar(foo, 10)
}
