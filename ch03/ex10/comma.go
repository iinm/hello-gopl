package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("123"))
	fmt.Println(comma("123456"))
	fmt.Println(comma("1234567"))
}

func comma(s string) string {
	var buf bytes.Buffer
	n := len(s)
	for i := 0; i < n; i++ {
		buf.WriteByte(s[i])
		m := n - i - 1 // 残りの文字数
		if m > 0 && m%3 == 0 {
			buf.WriteByte(',')
		}
	}
	return buf.String()
}
