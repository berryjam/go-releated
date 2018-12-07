package flowcontrol

import (
	"time"
	"fmt"
)

func TestSwitchWithNoCondition() {
	t := time.Now()
	/**
	This construct can be a clean way to write long if-then-else chains.
	switch可以不加条件，直接在case填入条件，可以替代很长的if-tehn-else代码，提高可读性
	 */
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
