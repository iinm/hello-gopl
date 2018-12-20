package sexpr

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"sort"
)

type Encoder struct {
	out io.Writer
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w}
}

func (enc *Encoder) Encode(v interface{}) error {
	if err := encode(enc.out, reflect.ValueOf(v), 0); err != nil {
		return err
	}
	return nil
}

type byteCountWriter struct {
	w         io.Writer
	ByteCount int
}

func (cw *byteCountWriter) Write(p []byte) (int, error) {
	n, err := cw.w.Write(p)
	cw.ByteCount += n
	return n, err
}

func encode(out io.Writer, v reflect.Value, indent int) error {
	countWriter := &byteCountWriter{out, 0}

	switch v.Kind() {
	case reflect.Invalid:
		fmt.Fprint(countWriter, "nil")

	case reflect.Bool:
		if v.Bool() {
			fmt.Fprint(countWriter, "t")
		} else {
			fmt.Fprint(countWriter, "nil")
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Fprintf(countWriter, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(countWriter, "%d", v.Uint())

	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(countWriter, "%g", v.Float())

	case reflect.String:
		fmt.Fprintf(countWriter, "%q", v.String())

	case reflect.Complex64, reflect.Complex128:
		fmt.Fprintf(countWriter, "#C(%g %g)", real(v.Complex()), imag(v.Complex()))

	case reflect.Ptr:
		return encode(countWriter, v.Elem(), 0)

	case reflect.Array, reflect.Slice: // (value ...)
		fmt.Fprint(countWriter, "(")
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				fmt.Fprintf(countWriter, "\n%*s", indent+1, "")
			}
			if err := encode(countWriter, v.Index(i), 0); err != nil {
				return err
			}
		}
		fmt.Fprint(countWriter, ")")

	case reflect.Struct: // ((name value) ...)
		fmt.Fprint(countWriter, "(")
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				fmt.Fprintf(countWriter, "\n%*s", indent+1, "")
			}
			fieldName := v.Type().Field(i).Name
			fmt.Fprintf(countWriter, "(%s ", fieldName)
			if err := encode(countWriter, v.Field(i), indent+len(fieldName)+3); err != nil {
				return err
			}
			fmt.Fprint(countWriter, ")")
		}
		fmt.Fprint(countWriter, ")")

	case reflect.Map: // ((key value) ...)
		// エンコード結果がおなじになるようにキーをソートしておく
		var keys []reflect.Value
		for _, key := range v.MapKeys() {
			keys = append(keys, key)
		}
		sort.Slice(keys, func(i, j int) bool {
			return stringify(keys[i]) < stringify(keys[j])
		})

		fmt.Fprint(countWriter, "(")
		for i, key := range keys {
			if i > 0 {
				fmt.Fprintf(countWriter, "\n%*s", indent+1, "")
			}
			fmt.Fprint(countWriter, "(")
			lenBeforeWrite := countWriter.ByteCount
			if err := encode(countWriter, key, 0); err != nil {
				return err
			}
			lenAfterWrite := countWriter.ByteCount
			fmt.Fprint(countWriter, " ")
			if err := encode(countWriter, v.MapIndex(key), indent+(lenAfterWrite-lenBeforeWrite)+3); err != nil {
				return err
			}
			fmt.Fprint(countWriter, ")")
		}
		fmt.Fprint(countWriter, ")")

	case reflect.Interface:
		fmt.Fprintf(countWriter, "(%q '", reflect.Indirect(v).Type())
		if err := encode(countWriter, reflect.Indirect(v).Elem(), 0); err != nil {
			return err
		}
		fmt.Fprint(countWriter, ")")

	default: // chan, func
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}

func stringify(v reflect.Value) string {
	var buf bytes.Buffer
	if err := encode(&buf, v, 0); err != nil {
		return ""
	}
	return buf.String()
}
