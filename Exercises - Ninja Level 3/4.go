package main

import "fmt"

func main() {
	years := 1999
	for {
		if years > 2021 {
			break
		}
		fmt.Println(years)
		years++
	}
}
