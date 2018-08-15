package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma("123"))
	fmt.Println(comma("1234567"))
	fmt.Println(comma("-123"))
	fmt.Println(comma("-1234567"))
	fmt.Println(comma("-123.1234"))
	fmt.Println(comma("+1234567.1234"))
}

func comma(s string) string {
	var buf bytes.Buffer
	// 符号があれば出力
	if len(s) > 0 && (s[0] == '-' || s[0] == '+') {
		buf.WriteByte(s[0])
		s = s[1:]
	}
	// 小数点があれば、最後に出力するために保持
	tail := ""
	if i := strings.Index(s, "."); i > -1 {
		s, tail = s[:i], s[i:]
	}

	n := len(s)
	for i := 0; i < n; i++ {
		buf.WriteByte(s[i])
		m := n - i - 1 // 残りの文字数
		if m > 0 && m%3 == 0 {
			buf.WriteByte(',')
		}
	}
	// 小数点以下の書き込み
	buf.WriteString(tail)

	return buf.String()
}
