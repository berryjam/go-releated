package flowcontrol

import (
	"fmt"
	"os"
	"io"
)

func c() (i int) {
	defer func() {
		i++
	}()
	return 1
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	// 此时src处于打开状态，defer后确保CopyFile函数退出前能够执行src.Close()
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	// 同src
	defer dst.Close()

	written, err = io.Copy(dst, src)
	if err != nil {
		return
	}
	return
}

func main() {

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		/**
		Deferred function calls are pushed onto a stack.
		 When a function returns, its deferred calls are executed in last-in-first-out orderapi.
		defer修饰的函数调用遵循"先进后出"原则
		 */
		defer fmt.Println(i)
	}

	fmt.Println("done")

	fmt.Print(c(), "\n")

	//CopyFile("/Users/berryjam/Downloads/tmp.go",
	//	"/Users/berryjam/Documents/workspace/go-releated/src/Basics/Flow control statements:for,if,else,switch and defer/defer-multi.go")
}
