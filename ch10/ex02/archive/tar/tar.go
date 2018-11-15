package tar

import (
	"archive/tar"
	"io"

	"../../archive"
)

func init() {
	archive.AddReader(ReadAll)
}

func ReadAll(in io.Reader) ([]*archive.Entry, error) {
	tr := tar.NewReader(in)
	var entries []*archive.Entry
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			return nil, err
		}
		entries = append(entries, &archive.Entry{Name: hdr.Name})
	}
	return entries, nil
}
