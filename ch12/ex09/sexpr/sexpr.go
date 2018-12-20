package sexpr

import (
	"fmt"
	"io"
	"reflect"
	"strconv"
	"text/scanner"
)

type Decoder struct {
	input io.Reader
	scan  scanner.Scanner
	token Token
	err   error
}

func NewDecoder(r io.Reader) *Decoder {
	dec := &Decoder{
		input: r,
		scan:  scanner.Scanner{Mode: scanner.GoTokens},
	}
	dec.scan.Init(dec.input)
	return dec
}

func (dec *Decoder) Decode(v interface{}) (err error) {
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("error at %s: %v", dec.scan.Position, x)
		}
	}()
	dec.next()
	dec.read(reflect.ValueOf(v).Elem())
	return nil
}

type Token interface{}
type Symbol string
type String string
type Int int64
type StartList struct{}
type EndList struct{}

func (dec *Decoder) read(v reflect.Value) {
	switch tok := dec.token.(type) {
	case Symbol:
		if tok == "nil" {
			v.Set(reflect.Zero(v.Type()))
			dec.next()
		}
	case String:
		v.SetString(string(tok))
		dec.next()
	case Int:
		v.SetInt(int64(tok))
		dec.next()
	case StartList:
		dec.consumeStartList()
		dec.readList(v)
		dec.consumeEndList()
	default:
		panic(fmt.Errorf("unexpected token: %v", tok))
	}
}

func (dec *Decoder) readList(v reflect.Value) {
	switch v.Kind() {
	case reflect.Array:
		for i := 0; !dec.atEndList(); i++ {
			dec.read(v.Index(i))
		}
	case reflect.Slice:
		for !dec.atEndList() {
			item := reflect.New(v.Type().Elem()).Elem()
			dec.read(item)
			v.Set(reflect.Append(v, item))
		}
	case reflect.Struct: // ((name value) ...)
		for !dec.atEndList() {
			dec.consumeStartList()
			switch tok := dec.token.(type) {
			case Symbol:
				name := string(tok)
				dec.next()
				dec.read(v.FieldByName(name))
			default:
				panic(fmt.Errorf("got token %v, want field name", tok))
			}
			dec.consumeEndList()
		}
	case reflect.Map:
		v.Set(reflect.MakeMap(v.Type()))
		for !dec.atEndList() {
			dec.consumeStartList()
			key := reflect.New(v.Type().Key()).Elem()
			dec.read(key)
			value := reflect.New(v.Type().Elem()).Elem()
			dec.read(value)
			v.SetMapIndex(key, value)
			dec.consumeEndList()
		}
	default:
		panic(fmt.Errorf("cannot decode list into %v", v.Type()))
	}
}

func (dec *Decoder) next() {
	dec.token, dec.err = dec.Token()
	if dec.err != io.EOF && dec.err != nil {
		panic(dec.err)
	}
}

func (dec *Decoder) Token() (Token, error) {
	r := dec.scan.Scan()
	switch r {
	case scanner.Ident:
		return Symbol(dec.scan.TokenText()), nil
	case scanner.String:
		s, _ := strconv.Unquote(dec.scan.TokenText())
		return String(s), nil
	case scanner.Int:
		i, _ := strconv.Atoi(dec.scan.TokenText())
		return Int(i), nil
	case '(':
		return StartList{}, nil
	case ')':
		return EndList{}, nil
	case scanner.EOF:
		return nil, io.EOF
	}
	return nil, fmt.Errorf("unexpected token %q", dec.scan.TokenText())
}

func (dec *Decoder) consumeStartList() {
	switch tok := dec.token.(type) {
	case StartList:
		dec.next()
	default:
		panic(fmt.Errorf("got %v, want StartList", tok))
	}
}

func (dec *Decoder) consumeEndList() {
	switch tok := dec.token.(type) {
	case EndList:
		dec.next()
	default:
		panic(fmt.Errorf("got %v, want EndList", tok))
	}
}

func (dec *Decoder) atEndList() bool {
	if dec.err == io.EOF {
		panic("end of file")
	}
	switch dec.token.(type) {
	case EndList:
		return true
	default:
		return false
	}
}
