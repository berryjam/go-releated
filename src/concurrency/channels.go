package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func selectChannel() {
	channelA := make(chan int)
	channelB := make(chan int)
	go func() {
		time.Sleep(5 * time.Second)
		channelA <- 1
	}()

	for {
		select {
		case a := <-channelA:
			fmt.Println(a)
			return
		case b := <-channelB:
			fmt.Println(b)
			return
		}
	}
}

func testChannelClose() {
	ch := make(chan int)
	go func() {
		time.Sleep(5 * time.Second)
		close(ch)
	}()

	// Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression: after
	v, ok := <-ch
	if !ok { // ok is false if there are no more values to receive and the channel is closed.
		fmt.Println("channel has been closed.")
		return
	}
	fmt.Println(v)
}

func main() {
	//s := []int{7, 2, 8, -9, 4, 0}
	//
	//c := make(chan int) // Like maps and slices, channels must be created before use:
	//go sum(s[:len(s)/2], c)
	//go sum(s[len(s)/2:], c)
	//x, y := <-c, <-c // receive from c
	//fmt.Println(x, y, x+y)

	//selectChannel()
	testChannelClose()
}
