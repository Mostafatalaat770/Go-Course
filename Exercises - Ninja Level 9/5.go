package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var counter int64

func fun1() {
	atomic.AddInt64(&counter, 1)
	runtime.Gosched()
	fmt.Println(atomic.LoadInt64(&counter))
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
