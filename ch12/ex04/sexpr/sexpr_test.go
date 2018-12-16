package sexpr_test

import (
	"fmt"

	sexpr "."
)

func Example_marshal() {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
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
	bs, _ := sexpr.Marshal(strangelove)
	fmt.Print(string(bs))

	// Output:
	// ((Title "Dr. Strangelove")
	//  (Subtitle "How I Learned to Stop Worrying and Love the Bomb")
	//  (Year 1964)
	//  (Actor (("Dr. Strangelove" "Peter Sellers")
	//          ("Grp. Capt. Lionel Mandrake" "Peter Sellers")
	//          ("Pres. Merkin Muffley" "Peter Sellers")
	//          ("Gen. Buck Turgidson" "George C. Scott")
	//          ("Brig. Gen. Jack D. Ripper" "Sterling Hayden")
	//          ("Maj. T.J. \"King\" Kong" "Slim Pickens")))
	//  (Oscars ("Best Actor (Nomin.)"
	//           "Best Adapted Screenplay (Nomin.)"
	//           "Best Director (Nomin.)"
	//           "Best Picture (Nomin.)"))
	//  (Sequel nil))
}

func Example_marshal_mapList() {
	v := map[string][]string{
		"Dr. Strangelove":            []string{"Peter Sellers", "hoge"},
		"Grp. Capt. Lionel Mandrake": []string{"Peter Sellers", "foo"},
		"Pres. Merkin Muffley":       []string{"Peter Sellers"},
	}
	bs, _ := sexpr.Marshal(v)
	fmt.Print(string(bs))

	// Output:
	// (("Pres. Merkin Muffley" ("Peter Sellers"))
	//  ("Dr. Strangelove" ("Peter Sellers"
	//                      "hoge"))
	//  ("Grp. Capt. Lionel Mandrake" ("Peter Sellers"
	//                                 "foo")))
}
