package main

import (
	"fmt"
	"net"
)

func main() {
	l, err := net.ListenPacket("udp", ":8889")
	if err != nil {
		panic(err)
	}
	fmt.Println("listen to 8889")
	for {
		buf := make([]byte, 1024)
		n, addr, err := l.ReadFrom(buf)
		if err != nil {
			fmt.Printf("read packet err:[%v]\n", err)
			continue
		} else {
			fmt.Printf("req:[%v] num:[%v] addr:[%v]\n", string(buf[:n]), n, addr)
		}
	}
}
