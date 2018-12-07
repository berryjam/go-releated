package main

import (
	"sync"
	"fmt"
)

type threadSafeSet struct {
	s []int
	sync.Mutex
}

func (set *threadSafeSet) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		set.Lock()

		for elem := range set.s {
			ch <- elem
		}

		close(ch)
		set.Unlock()
	}()
	return ch
}

func main() {
	set := threadSafeSet{s: make([]int, 0)}
	for i := 0; i < 10; i++ {
		set.s = append(set.s, 1)
	}
	ch := set.Iter()
	for c := range ch {
		fmt.Printf("%+v\n", c)
	}
}
