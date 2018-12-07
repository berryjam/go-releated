package methodsinterfaces

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
/**
Implement a Reader type that emits an infinite stream of the ASCII character 'A'.
 */
func (myReader MyReader) Read(data []byte) (int, error) {
	for index := range data {
		data[index] = 'A'
	}
	return len(data), nil
}

func main() {
	reader.Validate(MyReader{})
}
