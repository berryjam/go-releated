package main

import "fmt"

func main() {
	cha := make(chan bool)
	chb := make(chan bool)
	finish := make(chan bool)

	// print number
	go func(cha, chb chan bool) {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d%d", 2*i+1, 2*i+2)
			cha <- true
			<-chb
		}
	}(cha, chb)

	// print alphabet
	go func(cha, chb chan bool) {
		for i := 0; i < 5; i++ {
			<-cha
			fmt.Printf("%c%c", 'A'+2*i, 'A'+2*i+1)
			chb <- true
		}
		finish <- true
	}(cha, chb)

	<-finish
}
