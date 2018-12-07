package main

import (
	"bufio"
	"fmt"
	//"io/ioutil"
	//"io"
	"os"
	"github.com/cloudflare/cfssl/log"
	"runtime/pprof"
	"bytes"
	"strconv"
)

func parse(r *bufio.Reader) ([][]byte, error) {
	line, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	var argsCount int
	if line[0] == '*' {
		argsCount, _ = strconv.Atoi(string(line[1]))
	}
	args := make([][]byte, argsCount, argsCount)
	for i := 0; i < argsCount; i++ {
		if args[i], err = readArgument(r); err != nil {
			return nil, err
		}
	}
	return args, nil
}

func readArgument(r *bufio.Reader) ([]byte, error) {
	line, err := r.ReadString('\n')
	var argSize int
	_, err = fmt.Sscanf(line, "$%d\r", &argSize)
	if err != nil {
		return nil, err
	}

	data := make([]byte,argSize)
	_,err = r.Read(data)
	if err != nil {
		return nil,err
	}

	//data, err := ioutil.ReadAll(io.LimitReader(r, int64(argSize)))
	//if err != nil {
	//	return nil, err
	//}

	if len(data) != argSize {
		return nil, fmt.Errorf("error length of data.")
	}

	if b, err := r.ReadByte(); err != nil || b != '\r' {
		fmt.Printf("%s\n", string(b))
		return nil, fmt.Errorf("line should end with \\r\\n")
	}

	if b, err := r.ReadByte(); err != nil || b != '\n' {
		return nil, fmt.Errorf("line should end with \\r\\n")
	}

	return data, nil
}

func writeHeap(filename string) {
	f, err := os.Create(filename)

	if err != nil {
		log.Fatal(err)
	}
	pprof.WriteHeapProfile(f)

	defer f.Close()
}

func main() {
	s := []byte("*2\r\n$3\r\nGET\r\n$3\r\nKEY\r\n")
	for i := 0; i < 10000; i++ {
		buffer := bytes.NewReader(s)
		r := bufio.NewReaderSize(buffer, len(s))
		data, err := parse(r)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", data)
	}
	writeHeap("bytes.mprof.later")
}
