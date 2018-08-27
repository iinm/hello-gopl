package main

import "bytes"

func join(sep string, strs ...string) string {
	var buf bytes.Buffer
	for i, s := range strs {
		if i > 0 {
			buf.WriteString(sep)
		}
		buf.WriteString(s)
	}
	return buf.String()
}
