package main

import (
	"fmt"
	"net"
)

func main() {
	sendUDPWithoutProtocol()
}

func sendUDPWithoutProtocol() {
	data := []byte("这是一个数据包")
	conn, err := net.Dial("udp", "localhost:8889")
	if err != nil {
		fmt.Printf("dial udp fail,err:[%v]\n", err)
	} else {
		for i := 0; i < 1000; i++ {
			_, err := conn.Write(data)
			if err != nil {
				fmt.Printf("write failed , err: %v\n", err)
				break
			}
		}
	}
}
