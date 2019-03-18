package main

import (
	"fmt"
	"strconv"
	"os"
)

// http://redisdoc.com/topic/protocol.html 协议定义

// 解析请求协议
func ParseRequest(request []byte) string {
	if request == nil || len(request) < 4 {
		throwPanic("请求长度错误")
	}

	res := ""

	lines := make([][]byte, 0)

	lastStart := 0
	for idx := 0; idx < len(request); { // 以\r\n分割请求
		if request[idx] == '\r' { // 必须出现连续的\r\n
			if idx+1 >= len(request) || request[idx+1] != '\n' {
				throwPanic(`请求必须以\r\n作为分割符`)
			}
			lines = append(lines, request[lastStart:idx])
			lastStart = idx + 2
			idx += 2
		} else {
			idx++
		}
	}

	var argc, nextArgvBytes int
	var err error
	for idx, line := range lines {
		if idx == 0 { // first line
			if line[0] != '*' { // 首行第一个字节必须为*
				throwPanic("第一个字节必须为*")
			}
			argvStr := string(line[1:])
			argc, err = strconv.Atoi(argvStr)
			if err != nil {
				panic(err)
			}
			if len(lines) != 2*argc+1 { // 总行数为：参数数量*2+1
				throwPanic(fmt.Sprintf("参数数量不对，应该为:%v，实际为:%v", 2*argc+1, len(lines)))
			}
		} else {
			if idx%2 == 1 {
				if line[0] != '$' { // 参数字节数量行，以$打头
					throwPanic("参数字节数量行必须以$打头")
				}
				nextArgvBytes, err = strconv.Atoi(string(line[1:]))
				if err != nil {
					panic(err)
				}
			} else {
				if len(line) != nextArgvBytes {
					throwPanic(fmt.Sprintf("参数:'%v'字节数量错误，应该为:%v，实际为:%v", string(line), nextArgvBytes, len(line)))
				}
			}

		}
		res += string(line) + "\n"
	}

	return res
}

func ParseResponse(response []byte) string {
	
}

func throwPanic(s string) {
	panic(s)
}

func main() {
	args := os.Args
	exampleCmd := `./redis_protocol_parser request $'*3\r\n$3\r\nSET\r\n$5\r\nmykey\r\n$7\r\nmyvalue\r\n'`
	if len(args) != 3 {
		panic("命令格式:" + exampleCmd)
	}
	if args[1] == "request" {
		res := ParseRequest([]byte(args[2]))
		fmt.Printf("res:\n%+v\n", res)
	}
}
