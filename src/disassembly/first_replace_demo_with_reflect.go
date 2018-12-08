package main

import (
	"unsafe"
	"syscall"
	"reflect"
	"fmt"
)

func e() { fmt.Println("func e") }
func f() { fmt.Println("what the fuck") }

func JumpAssemblyDataWithReflect(p uintptr) []byte {
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
	//if reflect.ValueOf(from).Type() != reflect.ValueOf(to).Type() {
	//
	//}
	//if reflect.ValueOf(from).Kind() != reflect.Func {
	//
	//}

	fmt.Printf("reflect.ValueOf(to).Pointer():0x%x\n", reflect.ValueOf(to).Pointer())
	jumpAndExecToAssemblyData := JumpAssemblyDataWithReflect(reflect.ValueOf(to).Pointer())
	fmt.Printf("to:0x%x\n", reflect.ValueOf(to).Pointer())
	funcLocation := reflect.ValueOf(from).Pointer()
	fmt.Printf("from:0x%x\n", funcLocation)
	window := RawMemoryAccessWithReflect(funcLocation, len(jumpAndExecToAssemblyData))

	page := getPageWithReflect(funcLocation)
	syscall.Mprotect(page, syscall.PROT_READ|syscall.PROT_WRITE|syscall.PROT_EXEC)

	copy(window, jumpAndExecToAssemblyData)
}

func main() {
	replaceWithReflect(e, f)
	e()
}
