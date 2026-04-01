// https://go.dev/tour/methods/23

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13r *rot13Reader) Read(b []byte) (int, error) {
	n, err := rot13r.r.Read(b)

	for i, c := range b[:n] {
		if 'a' <= c && c <= 'z' {
			b[i] = (c-'a'+13)%26 + 'a'
		} else if 'A' <= c && c <= 'Z' {
			b[i] = (c-'A'+13)%26 + 'A'
		}
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
