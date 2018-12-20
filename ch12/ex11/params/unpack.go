package params

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

func Unpack(ptr interface{}) string {
	values := &url.Values{}
	v := reflect.ValueOf(ptr)
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}

		fieldValue := v.Field(i)
		switch fieldValue.Kind() {
		case reflect.Slice:
			for j := 0; j < fieldValue.Len(); j++ {
				elemValue := fieldValue.Index(j)
				values.Add(name, stringify(elemValue))
			}
		default:
			values.Add(name, stringify(fieldValue))
		}
	}
	return values.Encode()
}

func stringify(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Int:
		return fmt.Sprintf("%d", v.Int())
	case reflect.String:
		return v.String()
	case reflect.Bool:
		return fmt.Sprintf("%v", v.Bool())
	}
	return "" //todo
}
