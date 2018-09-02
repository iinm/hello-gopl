package counter

import (
	"unicode"
	"unicode/utf8"
)

type LineCounter struct {
	count      int
	writeCount int
}

func (c *LineCounter) Write(p []byte) (int, error) {
	for _, b := range p {
		if b == '\n' {
			c.count++
		}
	}
	c.writeCount++
	if c.writeCount == 1 && len(p) > 0 {
		c.count++
	}
	return len(p), nil
}

func (c *LineCounter) Count() int {
	return c.count
}

type WordCounter struct {
	count      int
	withinWord bool
}

func (c *WordCounter) Write(p []byte) (int, error) {
	for i, width := 0, 0; i < len(p); i += width {
		var r rune
		r, width = utf8.DecodeRune(p[i:])
		if !unicode.IsSpace(r) {
			c.withinWord = true
		} else {
			if c.withinWord {
				c.count++
			}
			c.withinWord = false
		}
	}
	return len(p), nil
}

func (c *WordCounter) Count() int {
	if c.withinWord {
		return c.count + 1
	}
	return c.count
}
