package sexpr_test

import (
	"bytes"
	"fmt"

	sexpr "."
)

func ExampleEncoder() {
	type Movie struct {
		Title    string            `sexpr:"title"`
		Subtitle string            `sexpr:"sub-title"`
		Year     int               `sexpr:"year"`
		Actor    map[string]string `sexpr:"actors"`
		Oscars   []string          `sexpr:"oscars"`
		Sequel   *string           `sexpr:"sequel"`
	}
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}

	buf := &bytes.Buffer{}
	encoder := sexpr.NewEncoder(buf)
	encoder.Encode(strangelove)
	fmt.Print(buf.String())

	// Output:
	// ((title "Dr. Strangelove")
	//  (sub-title "How I Learned to Stop Worrying and Love the Bomb")
	//  (year 1964)
	//  (actors (("Brig. Gen. Jack D. Ripper" "Sterling Hayden")
	//           ("Dr. Strangelove" "Peter Sellers")
	//           ("Gen. Buck Turgidson" "George C. Scott")
	//           ("Grp. Capt. Lionel Mandrake" "Peter Sellers")
	//           ("Maj. T.J. \"King\" Kong" "Slim Pickens")
	//           ("Pres. Merkin Muffley" "Peter Sellers")))
	//  (oscars ("Best Actor (Nomin.)"
	//           "Best Adapted Screenplay (Nomin.)"
	//           "Best Director (Nomin.)"
	//           "Best Picture (Nomin.)"))
	//  (sequel nil))
}
