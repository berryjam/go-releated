package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

func sendWithoutProtocol() {
	data := []byte("这是一个数据包")
	conn, err := net.DialTimeout("tcp", "localhost:8888", time.Second*30)
	if err != nil {
		fmt.Printf("connect failed, err : %v\n", err.Error())
		return
	}
	for i := 0; i < 1000; i++ {
		_, err = conn.Write(data)
		if err != nil {
			fmt.Printf("write failed , err : %v\n", err)
			break
		}
	}
}

func sendWithProtocol() {
	conn, err := net.DialTimeout("tcp", "localhost:8888", time.Second*30)
	if err != nil {
		fmt.Printf("connect failed, err : %v\n", err.Error())
		return
	}
	magicCodes := make([]byte, 2)
	magicCodes[0] = 0x12
	magicCodes[1] = 0x34
	body := []byte("这是一个数据包")
	bodyLenBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bodyLenBytes, uint32(len(body)))
	reqData := bytes.Join([][]byte{
		magicCodes,
		bodyLenBytes,
		body,
	}, []byte{})
	for i := 0; i < 1000; i++ {
		_, err = conn.Write(reqData)
		if err != nil {
			fmt.Printf("write failed , err : %v\n", err)
			break
		}
	}
}

func main() {
	//sendWithoutProtocol()
	sendWithProtocol()
}
