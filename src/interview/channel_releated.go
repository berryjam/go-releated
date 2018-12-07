package main

import "fmt"

func testSendChan() chan<- int {
	ch := make(chan<- int)
	return ch
}

func testReceiveChan(ch chan int) <-chan int {
	ch <- 100
	return ch
}

func main() {
	sendCh := testSendChan()
	//a := <- sendCh  the channel isn't able to receive value,the code will cause compile error
	sendCh <- 1

	ch := make(chan int)
	receiveCh := testReceiveChan(ch)
	//ch <- 100
	//receiveCh <- 1  the cahnnel isn't able to send value,the code will cause runtime error
	a := <-receiveCh
	fmt.Println(a)
}
