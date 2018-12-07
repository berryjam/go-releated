package flowcontrol

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i";
	}
	return fmt.Sprint(math.Sqrt(x))
}

func TestIf() {
	fmt.Println(sqrt(2), sqrt(-4))
}
