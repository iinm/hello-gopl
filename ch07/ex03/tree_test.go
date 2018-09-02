package tree

import (
	"fmt"
	"testing"
)

func TestTreeString(t *testing.T) {
	tests := []struct {
		root *tree
		want string
	}{
		{
			&tree{value: 0},
			`0
`,
		},
		{
			&tree{
				0,
				&tree{value: 1},
				&tree{value: 2},
			},
			`0
  1
  2
`,
		},
		{
			&tree{
				0,
				&tree{
					1,
					&tree{value: 3},
					&tree{value: 4},
				},
				&tree{
					2,
					&tree{value: 5},
					&tree{value: 6},
				},
			},
			`0
  1
    3
    4
  2
    5
    6
`,
		},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("%#v.String()", test.root)
		got := test.root.String()
		if got != test.want {
			t.Errorf("%s = %s want %s", descr, got, test.want)
		}
	}
}
