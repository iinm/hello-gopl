package main

import (
	"bytes"
	"unicode"
)

// s内の部分文字列"$foo"をf("foo")が返すテキストで置換する
func expand(s string, f func(string) string) string {
	const start = '$'
	const startL = len(string(start))
	var wordBuf, resultBuf bytes.Buffer

	flush := func() {
		if wordBuf.Len() == 0 {
			return
		}
		word := wordBuf.String()
		if word == string(start) {
			resultBuf.WriteString(word)
		} else {
			resultBuf.WriteString(f(word[startL:]))
		}
		wordBuf.Reset()
	}

	withinWord := false // 置換対象の単語を走査中か？
	for _, r := range s {
		if r == start { // 単語の開始
			flush()
			withinWord = true
		} else if unicode.IsSpace(r) { // 単語の終了
			flush()
			withinWord = false
		}

		if withinWord {
			wordBuf.WriteRune(r)
		} else {
			resultBuf.WriteRune(r)
		}
	}
	flush()

	return resultBuf.String()
}
