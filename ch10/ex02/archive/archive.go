package archive

import (
	"bytes"
	"errors"
	"io"
	"log"
)

type Entry struct {
	Name string
	// ...
}

type Reader func(io.Reader) ([]*Entry, error)

var readers []Reader

func AddReader(r Reader) {
	readers = append(readers, r)
}

func ReadAll(in io.Reader) ([]*Entry, error) {
	buf := &bytes.Buffer{}
	_, err := buf.ReadFrom(in) // todo: 大きなファイルを扱えない
	if err != nil {
		return nil, err
	}
	for _, reader := range readers {
		entries, err := reader(bytes.NewReader(buf.Bytes()))
		if err == nil {
			return entries, nil
		}
		log.Print(err)
	}
	// todo: サポートしていないのかその他のエラーかを区別したい
	return nil, errors.New("file type is not supported")
}
