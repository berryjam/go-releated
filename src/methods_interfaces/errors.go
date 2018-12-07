package methodsinterfaces

import (
	"time"
	"fmt"
	"strconv"
)

type MyError struct {
	When time.Time
	What string
}

/**
The error type is a built-in interface similar to fmt.Stringer
只要receiver类型实现了Error()，就表示这是一种error类型
 */
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	/**
	Functions often return an error value,
	and calling code should handle errors by testing whether the error equals nil.
	i, err := strconv.Atoi("42")
	if err != nil {
    		fmt.Printf("couldn't convert number: %v\n", err)
    		return
	}
	fmt.Println("Converted integer:", i)
	 */
	i, err := strconv.Atoi("24")
	if err != nil {
		fmt.Printf("couldn't convert number: %v\n", err)
	}
	fmt.Println("Converted integer:", i)
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
