package methodsinterfaces

import "fmt"

func main() {
	/**
	The interface type that specifies zero methods is known as the empty interface:
	and i can store value of any type
	 */
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

/**
An empty interface may hold values of any type
Empty interfaces are used by code that handles values of unknown type.
For example, fmt.Print takes any number of arguments of type interface{}.
 */
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}