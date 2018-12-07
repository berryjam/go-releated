package componets

import "fmt"

const Pi = 3.14

/**
 * Constants cannot be declared using the := syntax.
 * e.g const k = 4.2
 */


func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
