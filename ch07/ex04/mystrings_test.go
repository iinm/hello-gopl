package mystrings

import (
	"fmt"
	"io"
	"reflect"
	"testing"

	"golang.org/x/net/html"
)

func TestStringReader(t *testing.T) {
	tests := []struct {
		r         io.Reader
		dest      []byte
		wantedN   int
		wantedErr error
	}{
		{
			&StringReader{"", 0}, make([]byte, 0),
			0, io.EOF,
		},
		{
			&StringReader{"a", 0}, make([]byte, 0),
			0, nil,
		},
		{
			&StringReader{"foobar", 0}, make([]byte, 5),
			5, nil,
		},
		{
			&StringReader{"foobar", 0}, make([]byte, 6),
			6, io.EOF,
		},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("%v.Read(%v)", test.r, test.dest)
		n, err := test.r.Read(test.dest)
		if n != test.wantedN || err != test.wantedErr {
			t.Errorf("%s = (%d, %v), wanted (%d, %v)",
				descr, n, err, test.wantedN, test.wantedErr)
		}
	}
}

func TestNewReaderShouldReadAll(t *testing.T) {
	tests := []struct {
		src     string
		bufSize int
	}{
		{"hello", 3},
		{"hello", 5},
		{"hello", 6},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("NewReader(%q).Read()", test.src)

		r := NewReader(test.src)
		buf := make([]byte, test.bufSize)
		var read []byte
		for {
			n, err := r.Read(buf)
			if n == 0 {
				break
			}
			read = append(read, buf[:n]...)
			if err != nil {
				break
			}
		}

		if !reflect.DeepEqual(read, []byte(test.src)) {
			t.Errorf("%s -> read: %q (%[2]v), wanted: %q", descr, read, test.src)
		}
	}
}

func TestNewReaderWithHTMLParser(t *testing.T) {
	s := `<html><head></head><body><h1>h1</h1></body></html>`
	r := NewReader(s)
	doc, err := html.Parse(r)
	if err != nil {
		t.Error(err)
	}
	//     html ->    head ->    body ->     h1
	if doc.FirstChild.FirstChild.NextSibling.FirstChild.Data != "h1" {
		t.Error("could not parse correctly: h1 not found")
	}
}
