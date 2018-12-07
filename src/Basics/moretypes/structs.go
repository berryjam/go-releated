package moretypes

import "fmt"

/**
 A struct is a collection of fields.

 (And a type declaration does what you'd expect.)
 */

type StructsVertex struct {
	X int
	y int
}

func TestStructs() {
	fmt.Println(StructsVertex{1, 2})
	vertex := StructsVertex{1,2}
	vertex.X = 20
	vertex.y = 100
	fmt.Println(vertex)
}
