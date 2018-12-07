package moretypes

import "fmt"

/**
The capacity of a slice is the number of elements in the underlying array,
counting from the first element in the slice.
 */
func TestSliceLenCap() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.You can extend a slice's length by re-slicing it, provided it has sufficient capacity.
	s = s[:4]
	printSlice(s)

	// Drop ite first two values.
	s = s[2:]
	printSlice(s)

	s = s[0:4]    // s = s[0:5] panic: runtime error: slice bounds out of range
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
