package moretypes

import "fmt"

func TestSlices() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s[]int = primes[1:4] // slice的序列长度可变，[low, high)，遵循左闭右开原则，作用在数组上
	fmt.Println(s)
}
