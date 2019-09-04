package main

import "fmt"

func uleb128encode(num uint64) []byte {
	res := []byte{}

	if num == 0 {
		res = append(res, 0)
	} else {
		for num != 0 {
			b := (byte)(num & 0x7F)
			num >>= 7
			if num != 0 { /* more bytes to come */
				b |= 0x80
			}
			res = append(res, b)
		}
	}

	return res
}

func uleb128decode(bytes []byte) uint64 {
	if len(bytes) == 0 {
		panic("illegal input")
	}
	var res uint64 = 0
	var i uint8 = 0
	for {
		flag := bytes[i] & 0x80
		low7bit := bytes[i] & 0x7F
		res |= uint64(low7bit) << (7 * i)
		if flag != 0 {
			i++
		} else {
			break
		}
	}

	return res
}

func main() {
	encodedData := uleb128encode(12857)
	fmt.Printf("%x\n", encodedData)
	fmt.Printf("%+v\n", uleb128decode(encodedData))
}