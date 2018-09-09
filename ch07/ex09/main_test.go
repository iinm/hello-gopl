package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sort"
	"testing"
)

func TestHander(t *testing.T) {
	tests := []struct {
		query        string
		wantedStatus int
	}{
		{"", http.StatusOK},
		{"sort=", http.StatusOK},
		{"sort=title", http.StatusOK},
		{"sort=title,length", http.StatusOK},
		{"sort=nosuchkey", http.StatusBadRequest},
	}

	for _, test := range tests {
		url := "/?" + test.query
		descr := fmt.Sprintf("handler(%q)", url)

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handler)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != test.wantedStatus {
			t.Errorf("%s => status code: got %v want %v", descr, status, test.wantedStatus)
		}
	}
}

func TestStableSort(t *testing.T) {
	tests := []struct {
		descr  string
		tracks []*Track
		sort   func([]*Track)
		want   []*Track
	}{
		{
			descr: "do not sort",
			tracks: []*Track{
				{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
				{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
				{"Go", "Moby", "Moby", 1992, length("3m37s")},
				{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
			},
			sort: func(tracks []*Track) {
				byCustom := stableSort{t: tracks}
				sort.Sort(byCustom)
			},
			want: []*Track{
				{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
				{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
				{"Go", "Moby", "Moby", 1992, length("3m37s")},
				{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
			},
		},
		{
			descr: "sort by title, year",
			tracks: []*Track{
				{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
				{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
				{"Go", "Moby", "Moby", 1992, length("3m37s")},
				{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
			},
			sort: func(tracks []*Track) {
				byCustom := stableSort{t: tracks}
				byCustom.AddLessFunc(func(x, y *Track) bool { return x.Year < y.Year })
				byCustom.AddLessFunc(func(x, y *Track) bool { return x.Title < y.Title })
				sort.Sort(byCustom)
			},
			want: []*Track{
				{"Go", "Moby", "Moby", 1992, length("3m37s")},
				{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
				{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
				{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
			},
		},
		{
			descr: "sort by title, length (reversed)",
			tracks: []*Track{
				{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
				{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
				{"Go", "Moby", "Moby", 1992, length("3m37s")},
				{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
			},
			sort: func(tracks []*Track) {
				byCustom := stableSort{t: tracks}
				byCustom.AddLessFunc(func(x, y *Track) bool { return x.Length > y.Length })
				byCustom.AddLessFunc(func(x, y *Track) bool { return x.Title < y.Title })
				sort.Sort(byCustom)
			},
			want: []*Track{
				{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
				{"Go", "Moby", "Moby", 1992, length("3m37s")},
				{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
				{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
			},
		},
	}

	for _, test := range tests {
		test.sort(test.tracks)
		if !reflect.DeepEqual(test.tracks, test.want) {
			t.Errorf("%s ->\n%s,\nwant:\n%s",
				test.descr, stringify(test.tracks), stringify(test.want))
		}
	}
}

func stringify(tracks []*Track) string {
	buf := new(bytes.Buffer)
	printTracks(buf, tracks)
	return buf.String()
}
