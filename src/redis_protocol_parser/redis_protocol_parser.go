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
		throwPanic(fmt.Sprintf("请求长度错误，至少4个字节，实际字节数为:%v", len(request)))
	}

	res := ""

	lines := make([][]byte, 0)

	lastStart := 0
	for idx := 0; idx < len(request); { // 以\r\n分割请求
		if request[idx] == '\r' { // 必须出现连续的\r\n
			if idx+1 > len(request)-1 || request[idx+1] != '\n' {
				throwPanic(`请求必须以\r\n结尾，并作为分割符`)
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

// 解析回复协议
func ParseReply(reply []byte) string {
	var res string

	switch reply[0] {
	case '+':
		fallthrough
	case '-':
		if reply == nil || len(reply) < 3 {
			throwPanic(fmt.Sprintf("回复长度错误，至少为3个字节，实际为:%v\n", len(reply)))
		}
		res = parseStateReply(reply)
	case ':':
		if reply == nil || len(reply) < 3 {
			throwPanic(fmt.Sprintf("回复长度错误，至少为3个字节，实际为:%v\n", len(reply)))
		}
		res = parseIntegerReply(reply)
	case '$':
		res = parseBulkReply(reply)
	case '*':
		res = parseMultiBulkReply(reply)
	default:
		throwPanic("回复请求必须以+、-、:、$、*打头")
	}

	return res
}

func parseStateReply(reply []byte) string {
	for idx, b := range reply {
		if b == '\r' {
			if idx+1 != len(reply)-1 || reply[idx+1] != '\n' {
				throwPanic(`回复必须以\r\n结尾，并作为分割符`)
			}
			return string(reply[1:idx])
		}
	}

	return ""
}

func parseIntegerReply(reply []byte) string {
	for idx, b := range reply {
		if b == '\r' {
			if idx+1 < len(reply)-1 || reply[idx+1] != '\n' {
				throwPanic(`回复必须以\r\n结尾，并作为分割符`)
			}
			valStr := string(reply[1:idx])
			_, err := strconv.Atoi(valStr)
			if err != nil {
				panic(err)
			}
			return valStr
		}

	}
	return ""
}

func parseBulkReply(reply []byte) string {
	for idx, b := range reply {
		if b == '\r' {
			if idx+1 > len(reply)-1 || reply[idx+1] != '\n' {
				throwPanic(`回复必须以\r\n结尾，并作为分割符`)
			}
			lenStr := string(reply[1:idx])
			valLen, err := strconv.Atoi(lenStr)
			if err != nil {
				panic(err)
			}
			if valLen == -1 {
				return ""
			}

			if valLen > 512*1024*1024 {
				throwPanic("字符串的最大长度为 512 MB")
			}
			if idx+2+valLen != len(reply)-2 || reply[idx+2+valLen] != '\r' || idx+3+valLen != len(reply)-1 || reply[idx+3+valLen] != '\n' {
				throwPanic(`字符串长度不正确，回复最末尾必须以\r\n结尾`)
			}
			return string(reply[idx+2 : idx+2+valLen])
		}
	}
	return ""
}

func parseMultiBulkReply(reply []byte) string {
	if reply == nil || len(reply) < 4 {
		throwPanic(fmt.Sprintf("回复长度错误，至少4个字节，实际字节数为:%v", len(reply)))
	}

	res := ""

	lines := make([][]byte, 0)

	lastStart := 0
	for idx := 0; idx < len(reply); { // 以\r\n分割请求
		if reply[idx] == '\r' { // 必须出现连续的\r\n
			if idx+1 > len(reply)-1 || reply[idx+1] != '\n' {
				throwPanic(`请求必须以\r\n结尾，并作为分割符`)
			}
			lines = append(lines, reply[lastStart:idx+2])
			lastStart = idx + 2
			idx += 2
		} else {
			idx++
		}
	}

	var nextArgvBytes int
	var err error
	inMultiBulkReply := false
	for idx, line := range lines {
		if idx == 0 { // first line
			if line[0] != '*' { // 首行第一个字节必须为*
				throwPanic("第一个字节必须为*")
			}
			//argvStr := string(line[1:])
			//argc, err = strconv.Atoi(argvStr)
			//if err != nil {
			//	panic(err)
			//}
			//if len(lines) != 2*argc+1 { // 总行数为：参数数量*2+1
			//	throwPanic(fmt.Sprintf("参数数量不对，应该为:%v，实际为:%v", 2*argc+1, len(lines)))
			//}
			res += string(line[:len(line)-2]) + "\n"
		} else {
			if inMultiBulkReply {
				if nextArgvBytes+2 != len(line) {
					throwPanic(fmt.Sprintf("字符串长度错误，应该为:%v，实际为:%v\n", nextArgvBytes+2, len(line)))
				}
				res += string(line[:len(line)-2]) + "\n"
				inMultiBulkReply = false
			} else {
				switch line[0] {
				case '+':
				case '-':
					if len(line) < 3 {
						throwPanic(fmt.Sprintf("回复长度错误，至少为3个字节，实际为:%v\n", len(line)))
					}
					res += parseStateReply(line) + "\n"
					inMultiBulkReply = false
				case ':':
					if len(line) < 3 {
						throwPanic(fmt.Sprintf("回复长度错误，至少为3个字节，实际为:%v\n", len(line)))
					}
					res += parseIntegerReply(line) + "\n"
					inMultiBulkReply = false
				case '$':
					nextArgvBytes, err = strconv.Atoi(string(line[1 : len(line)-2]))
					if err != nil {
						panic(err)
					}
					//res += fmt.Sprintf("$%d", nextArgvBytes) + "\n"
					inMultiBulkReply = true
				}
			}
		}
	}

	return res
}

func throwPanic(s string) {
	panic(s)
}

func main() {
	args := os.Args
	exampleCmd := `./redis_protocol_parser request $'*3\r\n$3\r\nSET\r\n$5\r\nmykey\r\n$7\r\nmyvalue\r\n'\n./redis_protocol_parser reply $'*5\r\n:1\r\n:2\r\n:3\r\n:4\r\n$6\r\nfoobar\r\n'`
	if len(args) != 3 {
		panic("命令格式:" + exampleCmd)
	}
	if args[1] == "request" {
		res := ParseRequest([]byte(args[2]))
		fmt.Printf("res:\n%+v\n", res)
	} else if args[1] == "reply" {
		res := ParseReply([]byte(args[2]))
		fmt.Printf("res:\n%+v\n", res)
	}
}
