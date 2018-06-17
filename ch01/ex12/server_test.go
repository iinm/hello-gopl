package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLissajousHandler(t *testing.T) {
	tests := []struct {
		query        string
		wantedStatus int
	}{
		{"", http.StatusOK},
		{"cycles=6&res=0.0001&size=200&nframes=64&delay=4", http.StatusOK},
		{"cycles=foo", http.StatusBadRequest},
		{"size=-100", http.StatusInternalServerError},
	}

	for _, test := range tests {
		url := "/?" + test.query
		descr := fmt.Sprintf("lissajousHandler(%q)", url)

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(lissajousHandler)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != test.wantedStatus {
			t.Errorf("%s => status code: got %v want %v", descr, status, test.wantedStatus)
		}
	}
}
