package main

import (
	//"net"
	//"time"
	//"log"
	//"fmt"
	//"math"
	//"errors"
)

type SyntaxError struct {
	msg    string // error description
	Offset int64  // where the error occurred
}

func (e *SyntaxError) Error() string {
	return e.msg
}

//func Sqrt(f float64) (float64, error) {
//	if f < 0 {
//		return 0, errors.New("math: square root of negative number")
//	}
//
//	// implementation
//	return math.Sqrt(f), nil
//}

//func main() {
//
//	/**
//	if err := dec.Decode(&val); err != nil {
//    		if serr, ok := err.(*json.SyntaxError); ok {
//        		line, col := findLine(f, serr.Offset)
//        		return fmt.Errorf("%s:%d:%d: %v", f.Name(), line, col, err)
//    		}
//    		return err
//	}
//	 */
//	f, err := Sqrt(-1)
//
//	if nerr, ok := err.(net.Error); ok &&nerr.Temporary() {
//		time.Sleep(1e9)
//		// continue
//	}
//
//	if err != nil {
//		log.Fatal(err)
//	} else{
//		fmt.Printf("Sqrt(%d) = %d", -1, f)
//	}
//}

/**
func Decode() *SyntaxError {
	// error, which may lead to the caller's err != nil comparison to always be true.
	var err * SyntaxError // pre-declare error variable
	if an error condition {
		err = &SyntaxError{}
	}
	return err // error, err always equal non-nil, causes caller's err != nil comparison to always be true
}
 */

