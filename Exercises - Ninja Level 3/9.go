package main

import "fmt"

func main() {
	s := "favSport"
	switch s {
	case "Musty":
		fmt.Println("First Case")
	case "favsport":
		fmt.Println("Second Case")
	case "favSport":
		fmt.Println("Third Case")
	}
}
