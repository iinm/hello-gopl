package sexpr

import (
	"bytes"
	"reflect"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	type Movie struct {
		Title    string            `sexpr:"title"`
		Subtitle string            `sexpr:"subtitle"`
		Year     int               `sexpr:"year"`
		Actor    map[string]string `sexpr:"actors"`
		Oscars   []string
		Sequel   *string `sexpr:"sequel"`
	}
	input := []byte(`
	((title "Dr. Strangelove")
	 (subtitle "How I Learned to Stop Worrying and Love the Bomb")
	 (year 1964)
	 (actors (("Brig. Gen. Jack D. Ripper" "Sterling Hayden")
	          ("Dr. Strangelove" "Peter Sellers")
	          ("Gen. Buck Turgidson" "George C. Scott")
	          ("Grp. Capt. Lionel Mandrake" "Peter Sellers")
	          ("Maj. T.J. \"King\" Kong" "Slim Pickens")
	          ("Pres. Merkin Muffley" "Peter Sellers")))
	 (Oscars ("Best Actor (Nomin.)"
	          "Best Adapted Screenplay (Nomin.)"
	          "Best Director (Nomin.)"
	          "Best Picture (Nomin.)"))
	 (sequel nil))
	`)

	want := Movie{
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

	var got Movie
	//err := Unmarshal(input, &got)
	decoder := NewDecoder(bytes.NewReader(input))
	err := decoder.Decode(&got)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got:\n%#v\nwant:\n%#v", got, want)
	}
}
