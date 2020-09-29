package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type MyTimer struct {
	mu                   sync.Mutex
	lastRpcRevMillSecond int64
}

func (m *MyTimer) Rpc() {
	now := time.Now()
	fmt.Printf("call rpc now:%+v\n", now.UnixNano()/1e6)
	m.mu.Lock()
	m.lastRpcRevMillSecond = now.UnixNano() / 1e6
	m.mu.Unlock()
}

func main() {
	unstop := make(chan bool)
	myTimer := &MyTimer{
		lastRpcRevMillSecond: -1,
	}

	go func(m *MyTimer) {
		for {
			m.mu.Lock()
			if m.lastRpcRevMillSecond == -1 {
				m.mu.Unlock()
				fmt.Printf("now %+v first set timer,sleep 2s\n", time.Now().UnixNano()/1e6)
				time.Sleep(2000 * time.Millisecond)
			} else {
				now := time.Now()
				if m.lastRpcRevMillSecond+2000 <= now.UnixNano()/1e6 { // timeout
					fmt.Printf("timeout at %+v,reset timer,sleep 2s\n", now.UnixNano()/1e6)
					m.mu.Unlock()
					time.Sleep(2000 * time.Millisecond)
				} else {
					m.mu.Unlock()
					fmt.Printf("sleep %+v millisecond\n", m.lastRpcRevMillSecond+2000-now.UnixNano()/1e6)
					time.Sleep(time.Duration(m.lastRpcRevMillSecond+2000-now.UnixNano()/1e6) * time.Millisecond)
				}
			}
		}
	}(myTimer)

	go func(m *MyTimer) {
		for {
			r := rand.Int()
			if r%7 == 0 {
				m.Rpc()
			} else {
				fmt.Printf("no call rpc now:%+v\n", time.Now().UnixNano()/1e6)
			}
			time.Sleep(time.Second)
		}
	}(myTimer)

	<-unstop
}
