package main

import (
	"reflect"
	"fmt"
)

/**
Reflection in Go is used for determining information at runtime.
https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/02.6.html
 */

func main() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:",v.Float())
}
