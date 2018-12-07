package flowcontrol

import (
	"fmt"
	"runtime"
)

func TestSwitch() {
	fmt.Print("Go runs on ")
	/**
	 * A case body breaks automatically, unless it ends with a fallthrough statement
	 * switch的每个case会自动break
	 */
	switch os := runtime.GOOS; os{
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}
