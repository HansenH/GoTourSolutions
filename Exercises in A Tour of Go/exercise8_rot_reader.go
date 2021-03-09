package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader rot13Reader) Read(b []byte) (n int, err error) {
	n, err = reader.r.Read(b)
	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}
	return
}

func rot13(origin byte) byte {
	switch {
	case origin >= 65 && origin <= 90:
		return byte((origin-65+13)%26 + 65)
	case origin >= 97 && origin <= 122:
		return byte((origin-97+13)%26 + 97)
	default:
		return origin
	}
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
