package componets

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		x := rand.Intn(1000)
		fmt.Println(x)
	}
}
