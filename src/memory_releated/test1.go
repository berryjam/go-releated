package main

func foo() *int {
	x := 1
	return &x
}

func main() {
	x := foo()
	println(*x)
}

// go build -gcflags '-l' -o test1 test1.go
// go tool objdump -s "main\.foo" test1
