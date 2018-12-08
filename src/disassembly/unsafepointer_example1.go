package main

import (
	"fmt"
	"unsafe"
)

func ll() int { return 1 }

func main() {
	f := ll
	fmt.Printf("0x%x\n", *(*uintptr)(unsafe.Pointer(&f)))
}