package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func compressSpaces(b []byte) ([]byte, error) {
	i, j := 0, 0
	var previous []byte

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

		if previous == nil {
			previous = b[i : i+size]
			if isSpace(previous) {
				b[i] = ' '
				i, j = i+1, j+1
			} else {
				i += size
				j += size
			}
			continue
		}

		currentIsSpace := isSpace(b[i : i+size])
		if !(isSpace(previous) && currentIsSpace) {
			if currentIsSpace {
				b[j] = ' '
				j += 1
			} else {
				copy(b[j:j+size], b[i:i+size])
				j += size
			}
		}

		previous = b[i : i+size]
		i += size
	}

	return b[:j], nil
}

func isSpace(b []byte) bool {
	r, _ := utf8.DecodeRune(b)
	switch uint32(r) {
	case '\t', '\n', '\v', '\f', '\r', ' ', 0x85, 0xA0:
		return true
	}
	return false
}
