package main

import (
	"encoding/xml"
	"fmt"
	"reflect"
	"testing"
)

func TestParseMatcher(t *testing.T) {
	tests := []struct {
		s    string
		want ElementMatcher
	}{
		{"div", &TypeMatcher{"div"}},
		{".class", &ClassMatcher{"class"}},
		{"#id", &IdMatcher{"id"}},
		{
			"div.class",
			&CompositeMatcher{[]ElementMatcher{
				&TypeMatcher{"div"},
				&ClassMatcher{"class"},
			}},
		},
		{
			"div#id",
			&CompositeMatcher{[]ElementMatcher{
				&TypeMatcher{"div"},
				&IdMatcher{"id"},
			}},
		},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("ParseMatcher(%q)", test.s)
		m := ParseMatcher(test.s)
		if !reflect.DeepEqual(m, test.want) {
			t.Errorf("%s = %#v, want %#v", descr, m, test.want)
		}
	}
}

func TestElementMatcher(t *testing.T) {
	tests := []struct {
		matcher ElementMatcher
		el      *xml.StartElement
		want    bool
	}{
		// TypeMatcher
		{&TypeMatcher{"div"}, &xml.StartElement{Name: xml.Name{Local: "div"}}, true},
		{&TypeMatcher{"div"}, &xml.StartElement{Name: xml.Name{Local: "divv"}}, false},

		// IdMatcher
		{
			&IdMatcher{"foo"},
			&xml.StartElement{
				Name: xml.Name{Local: "div"},
				Attr: []xml.Attr{
					xml.Attr{Name: xml.Name{Local: "id"}, Value: "foo"},
				},
			},
			true,
		},
		{
			&IdMatcher{"foo"},
			&xml.StartElement{
				Name: xml.Name{Local: "div"},
				Attr: []xml.Attr{
					xml.Attr{Name: xml.Name{Local: "id"}, Value: "fooo"},
				},
			},
			false,
		},

		// ClassMatcher
		{
			&ClassMatcher{"foo"},
			&xml.StartElement{
				Name: xml.Name{Local: "div"},
				Attr: []xml.Attr{
					xml.Attr{Name: xml.Name{Local: "class"}, Value: "foo"},
				},
			},
			true,
		},
		{
			&ClassMatcher{"foo"},
			&xml.StartElement{
				Name: xml.Name{Local: "div"},
				Attr: []xml.Attr{
					xml.Attr{Name: xml.Name{Local: "class"}, Value: "fooo"},
				},
			},
			false,
		},

		// CompositeMatcher
		{
			&CompositeMatcher{[]ElementMatcher{
				&TypeMatcher{"div"},
				&IdMatcher{"foo"},
			}},
			&xml.StartElement{
				Name: xml.Name{Local: "div"},
				Attr: []xml.Attr{
					xml.Attr{Name: xml.Name{Local: "id"}, Value: "foo"},
				},
			},
			true,
		},
		{
			&CompositeMatcher{[]ElementMatcher{
				&TypeMatcher{"div"},
				&IdMatcher{"foo"},
			}},
			&xml.StartElement{
				Name: xml.Name{Local: "div"},
				Attr: []xml.Attr{
					xml.Attr{Name: xml.Name{Local: "id"}, Value: "fooo"},
				},
			},
			false,
		},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("%#v(%q)", test.matcher, test.el)
		got := test.matcher.Match(test.el)
		if got != test.want {
			t.Errorf("%s = %v, want %v", descr, got, test.want)
		}
	}
}
