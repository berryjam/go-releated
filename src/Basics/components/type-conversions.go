package componets

import (
	"math"
	"fmt"
)

func main() {
	var x, y = 3, 4
	var f = math.Sqrt(float64(x * x + y * y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
