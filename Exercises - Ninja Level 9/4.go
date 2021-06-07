package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mutex sync.Mutex
var counter int

func fun1() {
	mutex.Lock()
	x := counter
	x++
	counter = x
	mutex.Unlock()
	wg.Done()
}
func main() {
	wg.Add(5)
	go fun1()
	go fun1()
	go fun1()
	go fun1()
	go fun1()
	wg.Wait()
	fmt.Println(counter)

}
