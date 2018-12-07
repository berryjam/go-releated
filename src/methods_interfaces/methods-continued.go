package methodsinterfaces

import (
	"math"
	"fmt"
)

type MyFloat float64
type MyInt int

/**
You can only declare a method with a receiver whose type is defined in the same package as the method.
 You cannot declare a method with a receiver whose type is defined in another package (which includes the
 built-in types such as int)
 */
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (f MyInt) Abs() int {
	if f < 0 {
		return int(-f)
	}
	return int(f)
}

func main() {
	f1 := MyFloat(-math.Sqrt2)
	f2 := MyInt(-100)
	fmt.Println(f1.Abs())
	fmt.Println(f2.Abs())
}
