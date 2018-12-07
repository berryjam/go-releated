package flowcontrol

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%v >= %v\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func TestIfAndElse() {
	/** (Both calls to pow are executed and return before the call to fmt.Println in main begins.)
	 *  先执行两个pow函数，再执行main函数内部的fmt.Println，所以会先输出27 >= 20，再输出9 20
	 */
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
