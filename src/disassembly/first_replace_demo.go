package main

import (
	"fmt"
	"unsafe"
	"syscall"
)

func c() { fmt.Println("hello world") }
func d() { fmt.Println("world hello") }

func JumpAssemblyData(i func()) []byte {
	p := *(*uintptr)(unsafe.Pointer(&i))
	return []byte{
		0x48, 0xC7, 0xC2,
		byte(p),
		byte(p >> 8),
		byte(p >> 16),
		byte(p >> 24),
		//byte(p >> 32),
		//byte(p >> 40),
		//byte(p >> 48),
		//byte(p >> 56),
		0xFF, 0x22,
	}
}

func getPage(p uintptr) []byte {
	return (*(*[0xFFFFFF]byte)(unsafe.Pointer(p & ^uintptr(syscall.Getpagesize() - 1))))[:syscall.Getpagesize()]
}

func RawMemoryAccess(p uintptr) []byte {
	return (*(*[0xFF]byte)(unsafe.Pointer(p)))[:]
}

func replace(from, to func()) {
	jumpAndExecToAssemblyData := JumpAssemblyData(to)
	funcLocation := **(**uintptr)(unsafe.Pointer(&from))
	window := RawMemoryAccess(funcLocation)

	page := getPage(funcLocation)
	syscall.Mprotect(page, syscall.PROT_READ|syscall.PROT_WRITE|syscall.PROT_EXEC)

	copy(window, jumpAndExecToAssemblyData)
}

func main() {
	replace(c, d)
	c()
}
