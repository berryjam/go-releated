package moretypes

import "fmt"

func TestSlicesPointers() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"                // 改变slices的值也会改变引用的数组的值
	fmt.Println(a, b)
	fmt.Println(names)
}
