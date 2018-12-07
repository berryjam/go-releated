package main

import (
	"sync"
	"fmt"
)

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

func main() {
	ua := &UserAges{ages: make(map[string]int)}
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 20; i++ {
		go func(i int) {
			ua.Add(fmt.Sprintf("berry%d", i), i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	for i := 0; i < 20; i++ {
		fmt.Printf("name:berry%d age:%d\n", i, ua.Get(fmt.Sprintf("berry%d", i)))
	}
}
