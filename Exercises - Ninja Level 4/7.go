package main

import "fmt"

func main() {
	x := [][]string{{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "Helloooooo, James."},
	}

	for _, record := range x {
		fmt.Println(record)
		for _, data := range record {
			fmt.Println(data)
		}
	}
}
