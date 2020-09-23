package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type MyStruct struct {
	mu       sync.Mutex
	timerCh  chan bool
	rpcRecCh chan bool
}

func (m *MyStruct) Rpc() {
	fmt.Printf("call rpc\n")
	go func() {
		m.rpcRecCh <- true
	}()
}

func main() {
	unstop := make(chan bool)
	myStruct := &MyStruct{
		rpcRecCh: make(chan bool, 1),
	}

	go func(m *MyStruct) {
		timeoutCh := make(chan bool)
		timer := time.AfterFunc(2*time.Second, func() {
			timeoutCh <- true
		})
		for {
			select {
			case <-timeoutCh:
				fmt.Printf("timeout now:%+v\n", time.Now())
				timer = time.AfterFunc(2*time.Second, func() {
					timeoutCh <- true
				})
			case <-m.rpcRecCh:
				fmt.Printf("reset timer now:%+v\n", time.Now())
				timer.Stop()
				timer = time.AfterFunc(2*time.Second, func() {
					timeoutCh <- true
				})
			}
		}
	}(myStruct)

	go func(m *MyStruct) {
		for {
			r := rand.Int()
			if r%7 == 0 {
				m.Rpc()
				//break
			} else {
				fmt.Printf("no call rpc\n")
			}
			//m.Rpc()
			time.Sleep(1 * time.Second)
		}
	}(myStruct)

	<-unstop
}
