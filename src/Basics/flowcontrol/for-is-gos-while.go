package flowcontrol

import "fmt"

func TestFotIsGOsWhile() {
	sum := 1
	for sum < 1024 {
		sum += sum
	}
	fmt.Println(sum)
}
