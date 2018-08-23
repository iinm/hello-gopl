package main

import (
	"bytes"
	"unicode"
)

// s内の部分文字列"$foo"をf("foo")が返すテキストで置換する
func expand(s string, f func(string) string) string {
	var wordBuf, resultBuf bytes.Buffer
	withinWord := false // 置換対象の単語を走査中か？
	for _, r := range s {
		if r == '$' {
			if wordBuf.Len() > 0 {
				resultBuf.WriteString(f(wordBuf.String()))
				wordBuf.Reset()
			} else if withinWord {
				resultBuf.WriteRune('$')
			}
			withinWord = true
			continue // この文字自体は使わない
		} else if unicode.IsSpace(r) {
			// 空白文字は単語の終わりとする
			if wordBuf.Len() > 0 {
				resultBuf.WriteString(f(wordBuf.String()))
				wordBuf.Reset()
			} else if withinWord {
				resultBuf.WriteRune('$')
			}
			withinWord = false
		}

		if withinWord {
			wordBuf.WriteRune(r)
		} else {
			resultBuf.WriteRune(r)
		}
	}

	if wordBuf.Len() > 0 {
		resultBuf.WriteString(f(wordBuf.String()))
	} else if withinWord {
		resultBuf.WriteRune('$')
	}
	return resultBuf.String()
}
