package main

import (
	"fmt"
	"github.com/berryjam/monkey"
)

func testA() {
	fmt.Println("testA")
}

func testB() {
	fmt.Println("testB")
}

func main() {
	monkey.Patch(testA, testB)
	testA()
}
