package sexpr_test

import (
	"fmt"

	sexpr "."
)

func ExampleMarshal() {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
		Dubbing         map[string]string
		Music           []string
	}
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "",
		Year:     0,
		Actor:    nil,
		Oscars:   nil,
		Dubbing: map[string]string{
			"James": "Kimura",
			"Tom":   "Imanishi",
			"NoOne": "",
		},
		Music: []string{"Hello World!"},
	}
	bs, _ := sexpr.Marshal(strangelove)
	fmt.Print(string(bs))

	// Output:
	// ((Title "Dr. Strangelove")
	//  (Dubbing (("James" "Kimura")
	//            ("Tom" "Imanishi")))
	//  (Music ("Hello World!")))
}
