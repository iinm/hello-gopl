package zip

import (
	"archive/zip"
	"bytes"
	"io"

	"../../archive"
)

func init() {
	archive.AddReader(ReadAll)
}

func ReadAll(in io.Reader) ([]*archive.Entry, error) {
	buf := &bytes.Buffer{}
	buf.ReadFrom(in)
	r, err := zip.NewReader(bytes.NewReader(buf.Bytes()), int64(buf.Len()))
	if err != nil {
		return nil, err
	}

	var entries []*archive.Entry
	for _, f := range r.File {
		entries = append(entries, &archive.Entry{Name: f.Name})
	}

	return entries, nil
}
