package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	counter int32 = 0
)

func main() {
	wg := &sync.WaitGroup{}
	threadNum := 1000
	wg.Add(threadNum)

	for i := 0; i < threadNum; i++ {
		go incr(i, wg)
	}

	wg.Wait()
	fmt.Printf("counter:%+v\n", counter)
}

func incr(i int, wg *sync.WaitGroup) {
	defer wg.Done()

	spinNum := 0
	for {
		old := counter
		ok := atomic.CompareAndSwapInt32(&counter, old, old+1)
		if ok {
			break
		} else {
			spinNum++
		}
	}
	//counter++
	if spinNum != 0 {
		fmt.Printf("goroutine%+v: spinNum:%+v\n", i, spinNum)
	}
}
