package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic(err)
	}
	fmt.Println("listen to 8888")
	for {
		// 监听到新的连接，创建新的 goroutine 并处理
		conn, err := l.Accept()
		fmt.Printf("new conn from remote addr:[%+v]\n", conn.RemoteAddr())
		if err != nil {
			fmt.Printf("conn err:[%v]\n", err)
		} else {
			go handleTCP(conn)
			//go handleWithFrameBuilder(conn)
		}
	}
}

func handleTCP(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				continue
			} else {
				fmt.Printf("read connection content err:[%v]\n", err)
			}
		} else {
			req := bytes.NewBuffer(nil)
			reqNum, err := req.Write(buf[0:n])
			if err != nil {
				fmt.Printf("read req content err:[%v]\n", err)
			} else {
				fmt.Printf("req:[%v] num:[%v]\n", req.String(), reqNum)
			}
		}
	}
}

// 2 byte magic code + 4 byte body len + body data
func handleWithFrameBuilder(conn net.Conn) {
	defer conn.Close()
	for {
		magicCodes := make([]byte, 2)
		n, err := io.ReadFull(conn, magicCodes)
		if err != nil && err == io.EOF {
			continue
		}
		if err != nil || n != len(magicCodes) {
			fmt.Printf("read magic codes err:[%v] len:[%v]\n", err, n)
			continue
		} else if magicCodes[0] != 0x12 && magicCodes[1] != 0x34 {
			fmt.Printf("wrong magic code,magicCodes:[%v]\n", magicCodes)
			continue
		}
		bodyLenBytes := make([]byte, 4)
		n, err = io.ReadFull(conn, bodyLenBytes)
		if err != nil || n != len(bodyLenBytes) {
			fmt.Printf("read body len err:[%v] len:[%v]\n", err, n)
			continue
		}
		bodyLen := binary.BigEndian.Uint32(bodyLenBytes)
		body := make([]byte, bodyLen)
		n, err = io.ReadFull(conn, body)
		if err != nil || uint32(n) != bodyLen {
			fmt.Printf("read body err:[%v] len:[%v]\n", err, n)
		} else {
			fmt.Printf("req:[%v] num:[%v]\n", string(body), n)
		}
	}
}
