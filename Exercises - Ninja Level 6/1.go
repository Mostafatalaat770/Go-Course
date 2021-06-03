package main

import "fmt"

func foo() int {
	return 770
}
func bar() (int, string) {
	return 999, "Nubbie"
}
func main() {
	x, y := bar()
	fmt.Println(foo(), x, y)
}
