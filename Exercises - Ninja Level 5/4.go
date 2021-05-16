package main

import "fmt"

func main() {
	s := struct {
		first string
		age   int
	}{
		first: "Musty",
		age:   22,
	}

	fmt.Println(s.first)
	fmt.Println(s.age)
}
