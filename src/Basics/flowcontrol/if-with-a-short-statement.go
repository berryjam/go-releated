package flowcontrol

import (
	"fmt"
	"math"
)

func powIfWithAShortStatement(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func TestIfWithAShortStatement() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
