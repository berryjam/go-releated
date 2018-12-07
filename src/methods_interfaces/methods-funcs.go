package methodsinterfaces

import "fmt"

type Vertex struct {
	X, Y float64
}

/**
Remember: a method is just a function with a receiver argument.
Here's Abs written as a regular function with no change in functionality.
 */
func Abs(v Vertex) float64 {		// func (v Vertex) Abs() float64 is a method
	return v.X * v.X + v.Y * v.Y
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}
