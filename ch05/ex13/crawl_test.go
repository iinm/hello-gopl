package main

import (
	"fmt"
	"net/url"
	"testing"
)

func TestUrlToFilepath(t *testing.T) {
	tests := []struct {
		rawurl                        string
		wantedDirname, wantedFilename string
	}{
		{"http://host.domain", "host.domain", "index.html"},
		{"http://host.domain/", "host.domain", "index.html"},
		{"http://host.domain/path", "host.domain/path", "index.html"},
		{"http://host.domain/path/", "host.domain/path", "index.html"},
		{"http://host.domain/path?q=param", "host.domain/path", "index.html?q=param"},
		{"http://host.domain/path#section", "host.domain/path", "index.html#section"},
		{"http://host.domain/path.html?q=param", "host.domain", "path.html?q=param"},
		{"http://host.domain/path.html#section", "host.domain", "path.html#section"},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("urlToFilepath(%q)", test.rawurl)
		u, _ := url.Parse(test.rawurl)
		dir, file := urlToFilepath(u)
		if dir != test.wantedDirname || file != test.wantedFilename {
			t.Errorf("%s = (%q, %q), want (%q, %q)",
				descr, dir, file, test.wantedDirname, test.wantedFilename)
		}
	}
}
