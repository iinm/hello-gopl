package main

import "testing"

func TestModifyReturnValueInDefer(t *testing.T) {
	got := modifyReturnValueInDefer()
	if got != "modified!" {
		t.Errorf(`modifyReturnValueInDefer() = %q, want "modified!"`, got)
	}
}

func modifyReturnValueInDefer() (val string) { // 戻り値に名前を付けないとdeferで書き換えることができない
	val = "not modified"
	defer func() { val = "modified!" }()
	return val
}
