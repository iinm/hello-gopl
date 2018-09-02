package countingwriter

import "io"

type byteCounter struct {
	writer    io.Writer
	byteCount int64
}

func (c *byteCounter) Write(p []byte) (int, error) {
	n, err := c.writer.Write(p)
	c.byteCount += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	counter := byteCounter{w, 0}
	return &counter, &(counter.byteCount)
}
