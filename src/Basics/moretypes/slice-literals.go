package moretypes

import "fmt"

/** A slice literal is like an array literal without the length.
 *  array literal(没有名字，只是字面值): [3]bool{true, true, false}
 *  slice literal: []bool{true, true, false}
*/
func TestSliceLiterals() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
