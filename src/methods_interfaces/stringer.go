package methodsinterfaces

import "fmt"

type Person struct {
	Name string
	Age  int
}

/**
One of the most ubiquitous interfaces is Stringer defined by the fmt package.
The fmt package (and many others) look for this interface to print values.
类似java的toString方法
 */
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
