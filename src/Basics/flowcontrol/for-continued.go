package flowcontrol

import "fmt"

func TestForContinued() {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}
