package methodsinterfaces

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	/**
	Calling a method on a nil interface is a run-time error because there is no type
	inside the interface tuple to indicate which concrete method to call.
	 */
	i.M()
}

func describe(i I) {
	fmt.Println("(%v, %T)\n", i, i)
}
