package moretypes

import "fmt"

func TestRangeContinued() {
	pow := make([]int, 10)
	// If you only want the index, drop the ", value" entirely.省略value
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	// You can skip the index or value by assigning to _.，省略index需要加上_,
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
