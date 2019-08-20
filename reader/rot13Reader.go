package reader

import (
	"io"
)

type Rot13Reader struct {
	R io.Reader
}

func (r13Reader Rot13Reader) Read(b []byte) (num int, err error) {

	num, err = r13Reader.R.Read(b)

	for i := 0; i < len(b); i++ {
		char := b[i]
		if (char >= 'A' && char < 'N') || (char >= 'a' && char < 'n') {
			b[i] += 13
		} else if (char > 'M' && char <= 'Z') || (char > 'm' && char <= 'z') {
			b[i] -= 13
		}
	}

	return
}
