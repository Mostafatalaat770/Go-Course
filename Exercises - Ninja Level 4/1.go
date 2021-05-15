package main

import "fmt"

func main() {
	x := [5]int{}
	x[0] = 10
	x[1] = 20
	x[2] = 30
	x[3] = 40
	x[4] = 50
	for _, v := range x {
		fmt.Println(v)
	}

	fmt.Printf("%T\n", x)

}
