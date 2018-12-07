package moretypes

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	count := 0
	f1 := 0
	f2 := 1
	return func() int {
		if count == 0 {
			count ++
			return 0
		} else if count == 1 {
			count ++
			return 1
		} else {
			count ++
			f := f1 + f2
			f1 = f2
			f2 = f
			return f
		}
	}
}

func TestExerciseFibonacciClosure() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
