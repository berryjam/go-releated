package componets

import "fmt"

var i, j int = 1, 2

/**
 * Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
 * e.g k := 4  // error
 */

/**
 *Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
 */
func main() {
	var c, python, java = true, false, "no!"
	k := 4
	fmt.Println(i, j, c, python, java, k)
}
