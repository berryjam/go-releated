package moretypes

import "fmt"

/**
Go functions may be closures. A closure is a function value that references variables from outside its body.
The function may access and assign to the referenced variables; in this sense the function is "bound" to the variable

For example, the adder function returns a closure. Each closure is bound to its own sum variable.
 */
func adder() func(int) int {
	sum := 0	// 闭包，函数能够访问function body外的变量，这个函数相当于与该变量绑定在一起
	return func(x int) int {
		sum += x
		return sum
	}
}

func TestFunctionClosures() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2 * i))
	}
}
