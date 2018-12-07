package methodsinterfaces

import (
	"io"
	"strings"
	"os"
)

/**
A common pattern is an io.Reader that wraps another io.Reader, modifying the stream in some way
 */
type rot13Reader struct {
	r io.Reader
}

func (reader rot13Reader) Read(data []byte) (int, error) {
	n, err := reader.r.Read(data)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
