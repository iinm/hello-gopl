package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

func main() {
	byCustom := stableSort{t: tracks}
	//byCustom.AddLessFunc(func(x, y *Track) bool { return x.Year < y.Year })
	byCustom.AddLessFunc(func(x, y *Track) bool { return x.Length > y.Length })
	byCustom.AddLessFunc(func(x, y *Track) bool { return x.Title < y.Title })
	sort.Sort(byCustom)

	//sort.Sort(byYear(tracks))
	//sort.Stable(byTitle(tracks))

	printTracks(os.Stdout, tracks)
}

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
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

type byTitle []*Track

func (x byTitle) Len() int           { return len(x) }
func (x byTitle) Less(i, j int) bool { return x[i].Title < x[j].Title }
func (x byTitle) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type byYear []*Track

func (x byYear) Len() int           { return len(x) }
func (x byYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x byYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
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
