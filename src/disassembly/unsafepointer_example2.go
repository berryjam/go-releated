package main

import (
	"fmt"
	"unsafe"
)

func aa() int { return 1 }

func main() {
	f := aa
	fmt.Printf("0x%x\n", **(**uintptr)(unsafe.Pointer(&f)))
}