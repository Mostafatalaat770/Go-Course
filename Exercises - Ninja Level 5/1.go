package main

import "fmt"

type person struct {
	first       string
	last        string
	age         int
	favIceCream []string
}

func main() {
	p1 := person{
		first:       "Mostafa",
		last:        "Talaat",
		age:         22,
		favIceCream: []string{"Vanilla", "Pistachio", "Mango"},
	}
	p2 := person{
		first:       "Musty",
		last:        "Nubbie",
		age:         22,
		favIceCream: []string{"Vanilla", "Pistachio", "Mango", "Chocolate"},
	}
	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range p1.favIceCream {
		fmt.Println(i, v)
	}

	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for i, v := range p2.favIceCream {
		fmt.Println(i, v)
	}
}
