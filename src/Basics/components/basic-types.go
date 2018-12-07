package componets

import (
	"math/cmplx"
	"fmt"
)

/**
The example shows variables of several types,
and also that variable declarations may be "factored" into blocks,
as with import statements.
 */
var (
	ToBe bool = false
	MaxInt uint64 = 1 << 64 - 1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
}
