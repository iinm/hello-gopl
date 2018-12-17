package myjson

import (
	"bytes"
	"fmt"
	"reflect"
)

func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v), 0); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func encode(buf *bytes.Buffer, v reflect.Value, indent int) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("null")

	case reflect.Bool:
		if v.Bool() {
			buf.WriteString("true")
		} else {
			buf.WriteString("false")
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())

	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(buf, "%g", v.Float())

	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())

	case reflect.Complex64, reflect.Complex128:
		fmt.Fprintf(buf, `"#C(%g %g)"`, real(v.Complex()), imag(v.Complex()))

	case reflect.Ptr:
		return encode(buf, v.Elem(), 0)

	case reflect.Array, reflect.Slice: // (value ...)
		buf.WriteByte('[')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteString(fmt.Sprintf(",\n%*s", indent+1, ""))
			}
			if err := encode(buf, v.Index(i), 0); err != nil {
				return err
			}
		}
		buf.WriteByte(']')

	case reflect.Struct: // ((name value) ...)
		buf.WriteByte('{')
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				buf.WriteString(fmt.Sprintf(",\n%*s", indent+1, ""))
			}
			fieldName := v.Type().Field(i).Name
			fmt.Fprintf(buf, `"%s": `, fieldName)
			if err := encode(buf, v.Field(i), indent+len(fieldName)+5); err != nil {
				return err
			}
		}
		buf.WriteByte('}')

	case reflect.Map: // ((key value) ...)
		buf.WriteByte('{')
		for i, key := range v.MapKeys() {
			if i > 0 {
				buf.WriteString(fmt.Sprintf(",\n%*s", indent+1, ""))
			}
			lenBeforeWrite := buf.Len()
			if err := encode(buf, key, 0); err != nil {
				return err
			}
			buf.WriteString(": ")
			lenWrote := buf.Len() - lenBeforeWrite
			if err := encode(buf, v.MapIndex(key), indent+lenWrote+1); err != nil {
				return err
			}
		}
		buf.WriteByte('}')

	case reflect.Interface:
		fmt.Fprintf(buf, `"(%q '`, reflect.Indirect(v).Type())
		if err := encode(buf, reflect.Indirect(v).Elem(), 0); err != nil {
			return err
		}
		buf.WriteString(`)"`)

	default: // chan, func
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}
