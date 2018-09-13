package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

func main() {
	selector := ParseSelector(os.Args[1:])
	dec := xml.NewDecoder(os.Stdin)
	var stack []*xml.StartElement
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, &tok)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if selector.MatchAll(stack) {
				fmt.Printf("%s > %s\n", stringifyElemStack(stack), tok)
			}
		}
	}
}

type ElementMatcher interface {
	Match(*xml.StartElement) bool
}

type Selector struct {
	matchers []ElementMatcher
}

func (s *Selector) Add(m ElementMatcher) {
	s.matchers = append(s.matchers, m)
}

func (s *Selector) MatchAll(elems []*xml.StartElement) bool {
	matchers := s.matchers
	for len(matchers) <= len(elems) {
		if len(matchers) == 0 {
			return true
		}
		if matchers[0].Match(elems[0]) {
			matchers = matchers[1:]
		}
		elems = elems[1:]
	}
	return false
}

type TypeMatcher struct {
	typeName string
}

type IdMatcher struct {
	id string
}

type ClassMatcher struct {
	className string
}

type CompositeMatcher struct {
	matchers []ElementMatcher
}

func (s *TypeMatcher) Match(el *xml.StartElement) bool {
	return el.Name.Local == s.typeName
}

func (s *IdMatcher) Match(el *xml.StartElement) bool {
	for _, a := range el.Attr {
		if a.Name.Local == "id" && a.Value == s.id {
			return true
		}
	}
	return false
}

func (s *ClassMatcher) Match(el *xml.StartElement) bool {
	for _, a := range el.Attr {
		if a.Name.Local == "class" && a.Value == s.className {
			return true
		}
	}
	return false
}

func (s *CompositeMatcher) Match(el *xml.StartElement) bool {
	for _, matcher := range s.matchers {
		if !matcher.Match(el) {
			return false
		}
	}
	return true
}

func ParseSelector(ss []string) *Selector {
	selector := &Selector{}
	for _, s := range ss {
		matcher := ParseMatcher(s)
		selector.Add(matcher)
	}
	return selector
}

func ParseMatcher(s string) ElementMatcher {
	// type, #id, .class を分ける
	matcherStrs := []string{}
	buf := new(bytes.Buffer)
	for _, r := range s {
		if (r == '#' || r == '.') && buf.Len() > 0 {
			matcherStrs = append(matcherStrs, buf.String())
			buf.Reset()
		}
		buf.WriteRune(r)
	}
	if buf.Len() > 0 {
		matcherStrs = append(matcherStrs, buf.String())
		buf.Reset()
	}
	//fmt.Println(matcherStrs)

	matchers := []ElementMatcher{}
	for _, str := range matcherStrs {
		switch str[0] {
		case '#':
			matchers = append(matchers, &IdMatcher{str[1:]})
		case '.':
			matchers = append(matchers, &ClassMatcher{str[1:]})
		default:
			matchers = append(matchers, &TypeMatcher{str})
		}
	}

	if len(matchers) == 1 {
		return matchers[0]
	}
	return &CompositeMatcher{matchers}
}

func stringifyElemStack(elems []*xml.StartElement) string {
	buf := new(bytes.Buffer)
	for _, el := range elems {
		buf.WriteString(fmt.Sprintf(" > %s%s", el.Name.Local, stringifyAttrs(el.Attr)))
	}
	return buf.String()
}

func stringifyAttrs(attrs []xml.Attr) string {
	buf := new(bytes.Buffer)
	for _, a := range attrs {
		buf.WriteString(fmt.Sprintf(" %s=%q", a.Name.Local, a.Value))
	}
	return buf.String()
}
