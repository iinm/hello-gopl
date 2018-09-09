package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"sort"
	"strings"
	"text/tabwriter"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tracks := []*Track{
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
	}
	byCustom := &stableSort{t: tracks}

	if r.Form.Get(`sort`) != `` {
		keys := strings.Split(r.Form.Get(`sort`), ",")
		for i := len(keys) - 1; i >= 0; i-- {
			key := keys[i]
			if f, ok := lessFunctions[key]; ok {
				byCustom.AddLessFunc(f)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "invalid sort key: %q", key)
				return
			}
		}
	}

	sort.Sort(byCustom)
	printTracksHTML(w, tracks)
}

var lessFunctions = map[string]func(x, y *Track) bool{
	`title`:  func(x, y *Track) bool { return x.Title < y.Title },
	`artist`: func(x, y *Track) bool { return x.Artist < y.Artist },
	`album`:  func(x, y *Track) bool { return x.Album < y.Album },
	`year`:   func(x, y *Track) bool { return x.Year < y.Year },
	`length`: func(x, y *Track) bool { return x.Length < y.Length },
}

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

type stableSort struct {
	t             []*Track
	lessFunctions []func(x, y *Track) bool
}

func (x *stableSort) AddLessFunc(f func(x, y *Track) bool) {
	x.lessFunctions = append(x.lessFunctions, f)
}

func (x stableSort) Len() int      { return len(x.t) }
func (x stableSort) Swap(i, j int) { x.t[i], x.t[j] = x.t[j], x.t[i] }
func (x stableSort) Less(i, j int) bool {
	for k := len(x.lessFunctions) - 1; k >= 0; k-- {
		result := x.lessFunctions[k](x.t[i], x.t[j])
		inversedResult := x.lessFunctions[k](x.t[j], x.t[i])
		if result != inversedResult {
			return result
		}
	}
	return false
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

var htmlTemplate = template.Must(template.New("htmlTemplate").Parse(`
<table>
<tr style='text-align: left'>
  <th><a href="?sort=title">Title</a></th>
  <th><a href="?sort=artist">Artist</a></th>
  <th><a href="?sort=album">Album</a></th>
  <th><a href="?sort=year">Year</a></th>
  <th><a href="?sort=length">Length</a></th>
</tr>
{{range .}}
<tr>
  <td>{{.Title}}</td>
  <td>{{.Artist}}</td>
  <td>{{.Album}}</td>
  <td>{{.Year}}</td>
  <td>{{.Length}}</td>
</tr>
{{end}}
`))

func printTracksHTML(out io.Writer, tracks []*Track) {
	if err := htmlTemplate.Execute(out, tracks); err != nil {
		log.Fatal(err)
	}
}

func printTracks(out io.Writer, tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(out, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}
