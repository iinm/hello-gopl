package nodetree

import (
	"encoding/xml"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestMakeNodeTree(t *testing.T) {
	tests := []struct {
		xmlStr string
		want   *Element
	}{
		{``, nil},
		{
			`<html><body><h1 id="1">topic</h1></body></html>`,
			&Element{
				xml.Name{Space: "", Local: "html"},
				[]xml.Attr{},
				[]Node{
					&Element{
						xml.Name{Space: "", Local: "body"},
						[]xml.Attr{},
						[]Node{
							&Element{
								xml.Name{Space: "", Local: "h1"},
								[]xml.Attr{
									xml.Attr{Name: xml.Name{Space: "", Local: "id"}, Value: "1"},
								},
								[]Node{CharData("topic")},
							},
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("MakeNodeTree(%s)", test.xmlStr)
		root, err := MakeNodeTree(strings.NewReader(test.xmlStr))
		if err != nil || !reflect.DeepEqual(root, test.want) {
			t.Errorf("%s = (%s, %v),\nwant (%s, nil)", descr, root, err, test.want)
		}
	}
}
