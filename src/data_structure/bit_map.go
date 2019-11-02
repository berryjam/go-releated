package main

import "fmt"

type BitMap struct {
	bytes []byte
	nbits int
}

func MakeBitMap(n int) *BitMap {
	res := new(BitMap)
	res.nbits = n
	res.bytes = make([]byte, n/8+1)
	return res
}

func (b *BitMap) Set(k int) {
	if k > b.nbits {
		return
	}
	bytesIdx := k / 8
	bitIdx := uint8(k % 8)
	b.bytes[bytesIdx] |= 1 << bitIdx
}

func (b *BitMap) Get(k int) bool {
	if k > b.nbits {
		return false
	}
	bytesIdx := k / 8
	bitIdx := uint8(k % 8)
	return (b.bytes[bytesIdx] & (1 << bitIdx)) != 0
}

func main() {
	b := MakeBitMap(1000)
	fmt.Printf("%+v\n", b.Get(10))
	b.Set(10)
	fmt.Printf("%+v\n", b.Get(10))
}
