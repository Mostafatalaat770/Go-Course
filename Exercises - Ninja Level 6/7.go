package main

import (
	"fmt"
)

func main() {
	fun := func(x int) {
		fmt.Println(x)
	}

	fun(20)
}
