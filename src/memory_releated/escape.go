package main

import ()

func foo() *int {
	var x int
	return &x
}

func bar() int {
	x := new(int)
	*x = 1
	return *x
}

// go run -gcflags '-m -l' escape.go
func main() {}
