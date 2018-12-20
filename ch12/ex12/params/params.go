package params

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type Validator func(string) error

var validators = map[string]Validator{
	`email`: func(s string) error {
		if !strings.Contains(s, `@`) {
			return errors.New("email must contains @")
		}
		return nil
	},
}

func selectValidator(kind string) Validator {
	if v, ok := validators[kind]; ok {
		return v
	}
	return func(_ string) error { return nil }
}

func Unpack(req *http.Request, ptr interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}

	fields := make(map[string]reflect.Value)
	fieldValidators := make(map[string]Validator) // name -> validator
	v := reflect.ValueOf(ptr).Elem()              // the struct variable
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // a reflect.StructField
		tag := fieldInfo.Tag           // a reflect.StructTag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		fields[name] = v.Field(i)
		fieldValidators[name] = selectValidator(tag.Get("kind"))
	}

	for name, values := range req.Form {
		f := fields[name]
		validate := fieldValidators[name]
		if !f.IsValid() {
			continue
		}
		for _, value := range values {
			if err := validate(value); err != nil {
				return err
			}
			if f.Kind() == reflect.Slice {
				elem := reflect.New(f.Type().Elem()).Elem()
				if err := populate(elem, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				f.Set(reflect.Append(f, elem))
			} else {
				if err := populate(f, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}
		}
	}
	return nil
}

func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)

	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)

	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)

	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}
