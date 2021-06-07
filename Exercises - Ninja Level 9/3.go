package main

import (
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mutex sync.Mutex
var counter int

func fun1() {
	x := counter
	runtime.Gosched()
	x++
	counter = x
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

}
