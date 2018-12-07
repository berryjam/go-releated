package componets

import "fmt"

var c, python, java bool

/**
 * A var statement can be at package or function level. We see both in this example.
 */
func main() {
	var i int
	const f = "type=%T,value=%v"
	fmt.Printf(f + " " + f + " " + f + " " + f, i, i, c, c, python, python, java, java)
}
