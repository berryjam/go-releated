package methodsinterfaces

import "fmt"

func main() {
	var i interface{} = "hello"

	// A type assertion provides access to an interface value's underlying concrete value.
	s := i.(string)
	fmt.Println(s)

	/**
	Note the similarity between this syntax and that of reading from a map.
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
	 */
	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64)                // panic
	fmt.Println(f)
}
