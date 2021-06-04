package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func changeMe(p *person) {
	(*p).name = "Cherry woman, nubbiiee"
}

func main() {
	p := person{
		"Pizza man, nuuubu",
		22,
	}
	fmt.Println(p)
	changeMe(&p)
	fmt.Println(p)

}
