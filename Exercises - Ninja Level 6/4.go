package main

import "fmt"

type person struct {
	first  string
	second string
	age    int
}

func (p person) speak() {
	fmt.Println("My name is", p.first, p.second, ", I'm", p.age, "years old")
}

func main() {
	p := person{
		"Mostafa",
		"Talaat",
		22,
	}
	p.speak()
}
