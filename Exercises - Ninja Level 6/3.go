package main

import "fmt"

func foo() {

	fmt.Println("first function")

}

func main() {

	defer foo()
	func() {
		fmt.Println("second function")
	}()
}
