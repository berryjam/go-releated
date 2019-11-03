package main

import "fmt"

type BloomFilter struct {
	hashFuncs []func(n int) uint
	bytes     []byte
	nbits     int
}

func MakeBloomFilter(funcs []func(n int) uint, nbits int) *BloomFilter {
	res := new(BloomFilter)
	res.hashFuncs = funcs
	res.nbits = nbits
	res.bytes = make([]byte, nbits/8+1)
	return res
}

func (b *BloomFilter) Set(k int) {
	for _, fun := range b.hashFuncs {
		idx := fun(k)
		bytesIdx := idx / 8
		bitIdx := uint8(idx % 8)
		b.bytes[bytesIdx] |= 1 << bitIdx
	}
}

func (b *BloomFilter) Get(k int) bool {
	res := true

	for _, fun := range b.hashFuncs {
		idx := fun(k)
		bytesIdx := idx / 8
		bitIdx := uint8(idx % 8)
		if (b.bytes[bytesIdx] & (1 << bitIdx)) == 0 {
			return false
		}
	}

	return res
}

func main() {
	nbits := 1000
	b := MakeBloomFilter([]func(n int) uint{
		func(n int) uint {
			return uint(n % nbits)
		},
		func(n int) uint {
			return uint(n%nbits) + 1
		},
	}, nbits)
	fmt.Printf("%+v\n", b.Get(10))
	b.Set(10)
	fmt.Printf("%+v\n", b.Get(10))
}
