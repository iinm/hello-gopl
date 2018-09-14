package reader

import (
	"io"
)

type LimitedReader struct {
	src   io.Reader
	read  int
	limit int
}

func (r *LimitedReader) Read(p []byte) (n int, err error) {
	n, err = r.src.Read(p)
	if r.read+n > r.limit {
		// todo: 上でerrが発生してもEOF返してない？
		return r.limit - r.read, io.EOF
	}
	r.read += n
	return n, err
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitedReader{src: r, limit: int(n)}
}
