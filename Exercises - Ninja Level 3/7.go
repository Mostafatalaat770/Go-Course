package main

import "fmt"

func main() {
	x := 5
	if x == 5 {
		fmt.Println("Number is equals to 5!")
	} else if x%5 == 0 {
		fmt.Println("Divisble by 5!")
	} else {
		fmt.Println("Number is not Divisble by 5!")
	}
}
