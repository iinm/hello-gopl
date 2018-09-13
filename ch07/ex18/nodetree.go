package nodetree

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
)

type Node interface {
	String() string
} // CharData or *Element

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func MakeNodeTree(r io.Reader) (*Element, error) {
	dec := xml.NewDecoder(r)
	var root *Element
	stack := []*Element{}
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			el := &Element{tok.Name, tok.Attr, nil}
			if len(stack) == 0 {
				root = el
			} else {
				parent := stack[len(stack)-1]
				parent.Children = append(parent.Children, el)
			}
			stack = append(stack, el) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			el := CharData(tok)
			parent := stack[len(stack)-1]
			parent.Children = append(parent.Children, el)
		}
	}
	return root, nil
}

func (e *Element) String() string {
	buf := new(bytes.Buffer)
	for i, a := range e.Attr {
		if i > 0 {
			buf.WriteByte(' ')
		}
		buf.WriteString(fmt.Sprintf(`%s="%s"`, a.Name.Local, a.Value))
	}
	attrStr := buf.String()

	buf.Reset()
	for _, c := range e.Children {
		buf.WriteString(c.String())
	}
	childrenStr := buf.String()

	return fmt.Sprintf("<%s %s>%s</%[1]s>", e.Type.Local, attrStr, childrenStr)
}

func (c CharData) String() string { return string(c) }
