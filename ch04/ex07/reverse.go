package main

import (
	"errors"
	"fmt"
)

func reverseUnicodeBytes(b []byte) ([]byte, error) {
	i := 0
	for i < len(b) {
		var size int
		switch {
		case b[i]&(1<<7) == 0:
			size = 1
		case b[i]&(1<<7) != 0 && b[i]&(1<<6) != 0 && b[i]&(1<<5) == 0:
			size = 2
		case b[i]&(1<<7) != 0 && b[i]&(1<<6) != 0 && b[i]&(1<<5) != 0 && b[i]&(1<<4) == 0:
			size = 3
		case b[i]&(1<<7) != 0 && b[i]&(1<<6) != 0 && b[i]&(1<<5) != 0 && b[i]&(1<<4) != 0 && b[i]&(1<<3) == 0:
			size = 4
		default:
			return nil, errors.New(fmt.Sprintf("invalid code: %x %[1]b\n", b[i]))
		}

		// 文字単位でreverseする
		if size > 1 {
			reverse(b[i : i+size])
		}

		i += size
	}

	// 全体をreverseする
	reverse(b)
	return b, nil
}

func reverse(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}
