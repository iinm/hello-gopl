package myjson_test

import (
	"encoding/json"
	"testing"

	myjson "."
)

func TestMarshal(t *testing.T) {
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
	bs, err := myjson.Marshal(strangelove)
	if err != nil {
		t.Errorf("cannot marshal: %v", err)
	}
	t.Logf("marshal result:\n%s", string(bs))

	var unmarshaled Movie
	err = json.Unmarshal(bs, &unmarshaled)
	if err != nil {
		t.Errorf("cannot unmarshal: %v", err)
	}
	t.Logf("unmarshal result:\n%#v", unmarshaled)
}
