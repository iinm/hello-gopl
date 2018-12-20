package sexpr

import (
	"bytes"
	"reflect"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Score           float64
		Actor           map[string]string
		Oscars          []string
		Published       bool
		Sequel          *string
		Addon           interface{}
	}

	input := []byte(`
	((Title "Dr. Strangelove")
	 (Subtitle "How I Learned to Stop Worrying and Love the Bomb")
	 (Year 1964)
	 (Score 4.3)
	 (Actor (("Brig. Gen. Jack D. Ripper" "Sterling Hayden")
	         ("Dr. Strangelove" "Peter Sellers")
	         ("Gen. Buck Turgidson" "George C. Scott")
	         ("Grp. Capt. Lionel Mandrake" "Peter Sellers")
	         ("Maj. T.J. \"King\" Kong" "Slim Pickens")
	         ("Pres. Merkin Muffley" "Peter Sellers")))
	 (Oscars ("Best Actor (Nomin.)"
	          "Best Adapted Screenplay (Nomin.)"
	          "Best Director (Nomin.)"
	          "Best Picture (Nomin.)"))
	 (Published t)
	 (Addon t)
	 (Sequel nil))
	`)

	want := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Score:    4.3,
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
		Published: true,
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
