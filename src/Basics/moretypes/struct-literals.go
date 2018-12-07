package moretypes

import "fmt"

type StructLiteralsVertex struct {
	X, Y int
}

/**
  A struct literal denotes a newly allocated struct value by listing the values of its fields.
 */

var (
	v1 = StructLiteralsVertex{1, 2}        // has type Vertex
	v2 = StructLiteralsVertex{X:1}        // Y:0 is implicit
	v3 = StructLiteralsVertex{}                // X:0 and Y:0
	p = &StructLiteralsVertex{1, 2}        // has type *Vertex
)

func TestStructLiterals() {
	fmt.Println(v1, p, v2, v3)
}
