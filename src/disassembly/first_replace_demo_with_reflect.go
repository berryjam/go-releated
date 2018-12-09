package main

import (
	"unsafe"
	"syscall"
	"reflect"
	"fmt"
)

func e() float64 {
	fmt.Println("func e")
	return 0
}

func f() float64 {
	fmt.Println("func f")
	return 3.14
}

func g() string {
	return "return by g()"
}

func h() string {
	return "return by h()"
}

func JumpAssemblyDataWithReflect(p uintptr) []byte {
	return []byte{
		0x48, 0xC7, 0xC2,
		byte(p),
		byte(p >> 8),
		byte(p >> 16),
		byte(p >> 24),
		0xFF, 0xe2,
	}
}

func getPageWithReflect(p uintptr) []byte {
	return (*(*[0xFFFFFF]byte)(unsafe.Pointer(p & ^uintptr(syscall.Getpagesize() - 1))))[:syscall.Getpagesize()]
}

func RawMemoryAccessWithReflect(p uintptr, l int) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: p,
		Len:  l,
		Cap:  l,
	}))
	//return (*(*[0xFF]byte)(unsafe.Pointer(p)))[:]
}

func replaceWithReflect(from, to interface{}) {

	jumpAndExecToAssemblyData := JumpAssemblyDataWithReflect(reflect.ValueOf(to).Pointer())
	funcLocation := reflect.ValueOf(from).Pointer()
	window := RawMemoryAccessWithReflect(funcLocation, len(jumpAndExecToAssemblyData))

	page := getPageWithReflect(funcLocation)
	syscall.Mprotect(page, syscall.PROT_READ|syscall.PROT_WRITE|syscall.PROT_EXEC)

	copy(window, jumpAndExecToAssemblyData)
}

func main() {
	replaceWithReflect(e, f)
	fmt.Println(e())
	replaceWithReflect(g, h)
	fmt.Println(g())
}
