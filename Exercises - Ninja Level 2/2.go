package main

import "fmt"

func main() {
	g := 50 == 50
	h := 50 <= 49
	i := 50 >= 49
	j := 50 != 49
	k := 50 < 49
	l := 50 > 49

	fmt.Println(g, h, i, j, k, l)
}
