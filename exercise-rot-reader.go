package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r13 rot13Reader) Read(b []byte) (n int, err error) {
	n, err = r13.r.Read(b)

	if err != io.EOF {
		for i := 0; i < n; i++ {

			v := b[i]

			if (v >= uint8('A')) && (v <= uint8('M')) || (v >= uint8('a')) && (v <= uint8('m')) {
				v = v + 13
				b[i] = v
			} else if (v >= uint8('N')) && (v <= uint8('Z')) || (v >= uint8('n')) && (v <= uint8('z')) {
				v = v - 13
				b[i] = v
			}
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
