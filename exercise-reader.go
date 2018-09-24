package main

import (
	"golang.org/x/tour/reader"
	"io"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (mr MyReader) Read(b []byte) (n int, err error) {
	if len(b) != 0 {
		for i := range b {
			b[i] = 'A'
		}
		return len(b), nil
	}
	return 0, io.EOF
}

func main() {
	reader.Validate(MyReader{})
}
