package moretypes

import "fmt"

func TestMakingSlices() {
	/**
	Slices can be created with the built-in make function;
	this is how you create dynamically-sized arrays.
	 */
	a := make([]int, 5)        // len(a)=5, cap(a)=5
	printSliceMakingSlices("a", a)

	b := make([]int, 0, 5)        // len(b)=0, cap(b)=5
	printSliceMakingSlices("b", b)

	c := b[:2]
	printSliceMakingSlices("c", c)

	d := c[2:5]
	printSliceMakingSlices("d", d)
}

func printSliceMakingSlices(s string, x []int) {
	fmt.Println("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
