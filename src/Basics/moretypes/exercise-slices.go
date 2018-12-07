package moretypes

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	result := make([][] uint8, dx)
	for i := range result {
		result[i] = make([] uint8, dy)
		for j := range result[i] {
			//result[i][j] = uint8((i + j) / 2)
			// result[i][j] = uint8(i*j)
			result[i][j] = uint8(i ^ j)
		}
	}
	return result
}

func TestExerciseSlices() {
	pic.Show(Pic)
}
