package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	x1 := x[:5]
	x2 := x[5:]
	x3 := x[2:8]
	x4 := x[1:7]

	fmt.Println(x1, x2, x3, x4)
}
