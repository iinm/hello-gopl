package countingwriter

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	tests := []struct {
		inputs []string
	}{
		{[]string{""}},
		{[]string{"Hello World!"}},
		{[]string{"Hello World!", "こんにちは　世界！"}},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("CountingWriter.Write(%#v)", test.inputs)

		buf := new(bytes.Buffer)
		writer, byteCount := CountingWriter(buf)
		for _, s := range test.inputs {
			writer.Write([]byte(s))
		}

		wantedString := strings.Join(test.inputs, "")

		// bufに書き込めていること
		if buf.String() != wantedString {
			t.Errorf("%s -> buf.String() = %s, want %s", descr, buf.String(), wantedString)
		}

		// byte数をカウントできていること
		if *byteCount != int64(len(wantedString)) {
			t.Errorf("%s -> byteCount = %d, want %d", descr, *byteCount, len(wantedString))
		}
	}
}
