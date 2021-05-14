package main

import "fmt"

type Musty int

var x Musty
var y int

func main() {
	fmt.Printf("value of x: %v\ttype of x: %T\n", x, x)
	x = 42
	fmt.Println("value of x: ", x)

	y = int(x)
	fmt.Printf("value of x: %v\ttype of x: %T\n", y, y)
}
