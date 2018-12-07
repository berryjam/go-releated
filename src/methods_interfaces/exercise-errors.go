package methodsinterfaces

import (
	"fmt"
	"basics/flowcontrol"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	/**
	Note: a call to fmt.Sprint(e) inside the Error method will send the program into an infinite loop.
	You can avoid this by converting e first: fmt.Sprint(float64(e))
	因为 fmt.Sprintf(e)会递归调用e的Error方法，Error里面还是fmt.Sprintf(e)，导致无限递归
	 */
	//return fmt.Sprintf("cannot Sqrt negative number: %v", e)
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	} else {
		return flowcontrol.Sqrt(x), nil
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
