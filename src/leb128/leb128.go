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

func sleb128encode(value int64) []byte {
	res := []byte{}

	more := 1

	for more != 0 {
		b := (byte)(value & 0x7F)
		signFlag := (byte)(value & 0x40)
		value >>= 7
		if (value == 0 && signFlag == 0) || (value == -1 && signFlag != 0) {
			more = 0
		} else {
			b |= 0x80
		}
		res = append(res, b)
	}

	return res
}

func sleb128decode(bytes []byte) int64 {
	if len(bytes) == 0 {
		panic("illegal input")
	}
	var res uint64 = 0
	var i uint8 = 0
	isNegative := false
	var shift uint64 = 0
	for {
		flag := bytes[i] & 0x80
		low7bit := bytes[i] & 0x7F
		res |= uint64(low7bit) << (shift)
		shift+=7
		if flag != 0 {
			i++
		} else {
			signFlag := bytes[i] & 0x40
			if signFlag != 0 {
				isNegative = true
			}
			break
		}
	}
	if !isNegative {
		return int64(res)
	} else {
		tmp := int64(res)
		tmp |= -(1 << shift)
		return tmp
	}
}

func TestUleb128() {
	fmt.Println("TestUleb128")
	encodedData := uleb128encode(12857)
	fmt.Printf("%x\n", encodedData)
	fmt.Printf("%+v\n", uleb128decode(encodedData))
	fmt.Println()
}

func TestSleb128() {
	fmt.Println("TestSleb128")
	encodedData := uleb128encode(16256)
	fmt.Printf("%x\n", encodedData)
	fmt.Printf("%+v\n", uleb128decode(encodedData))

	encodedData = sleb128encode(-128)
	fmt.Printf("%x\n", encodedData)
	fmt.Printf("%+v\n", sleb128decode(encodedData))
	fmt.Println()
}

func TestSleb1281() {
	fmt.Println("TestSleb1281")
	encodedData := uleb128encode(12726)
	fmt.Printf("%x\n", encodedData)
	fmt.Printf("%+v\n", uleb128decode(encodedData))

	encodedData = sleb128encode(-3658)
	fmt.Printf("%x\n", encodedData)
	fmt.Printf("%+v\n", sleb128decode(encodedData))
	fmt.Println()
}

func main() {
	TestUleb128()

	TestSleb128()

	TestSleb1281()
}
