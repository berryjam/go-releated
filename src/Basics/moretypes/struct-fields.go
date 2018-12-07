package moretypes

import "fmt"

type StructFieldsVertex struct {
	X int
	Y int
}

func TestStructFields() {
	v := StructFieldsVertex{1, 2}
	// Struct fields are accessed using a dot.
	v.X = 4
	fmt.Println(v.X)
}
