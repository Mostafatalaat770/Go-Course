package main

import "fmt"

func foo(x ...int) int {

	var total int

	for _, v := range x {
		total += v
	}

	return total

}

func bar(x []int) int {
	var total int

	for _, v := range x {
		total += v
	}

	return total
}

func main() {
	x := []int{0, 1, 2, 3, 4, 5}

	fmt.Println(foo(x...), bar(x))
}
