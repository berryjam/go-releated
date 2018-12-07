package moretypes

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func TestRange() {
	/**
	When ranging over a slice, two values are returned for each iteration.
	The first is the index, and the second is a copy of the element at that index
	 */
	for i, v := range pow {
		// range每次迭代返回2个值，第1个为index，第2个为index在slice中的值
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
