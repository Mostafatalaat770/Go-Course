package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p *person) speak() {
	fmt.Println(p.first+" "+p.last, "is speaking.")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{"Mostafa", "Talaat", 22}
	saySomething(&p)
	saySomething(p) // does not work
}
