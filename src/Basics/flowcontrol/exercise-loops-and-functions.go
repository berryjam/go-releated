package flowcontrol

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	tmp := float64(z - ((z * z) - x) / (2 * z))
	for ; math.Abs(tmp - z) > 0.00001; {
		z = tmp
		tmp = z - ((z * z) - x) / (2 * z)
	}
	return tmp
}

func TestExerciseLoopsAndFunctions() {
	for i := 1; i < 11; i++ {
		fmt.Printf("Newton's Sqrt(%v)=%g\n", i, Sqrt(float64(i)))
		fmt.Printf("Math's Sqrt(%v)=%g\n\n", i, math.Sqrt(float64(i)))
	}

}
