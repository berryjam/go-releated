package methodsinterfaces

import (
	"strings"
	"fmt"
	"io"
)

func main() {
	r := strings.NewReader("Hello,Reader!")

	b := make([]byte, 8)
	for {
		// n:返回读取的字节数 err:返回错误
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
