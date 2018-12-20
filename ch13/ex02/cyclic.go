package cyclic

import (
	"reflect"
	"unsafe"
)

func cyclic(i interface{}) bool {
	v := reflect.ValueOf(i)
	seen := make(map[unsafe.Pointer]bool)
	return visit(v, seen)
}

func visit(v reflect.Value, seen map[unsafe.Pointer]bool) bool {
	if !v.IsValid() {
		return false
	}
	if v.CanAddr() {
		vptr := unsafe.Pointer(v.UnsafeAddr())
		if seen[vptr] {
			return true
		}
		seen[vptr] = true
	}
	switch v.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			return visit(v.Index(i), seen)
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			return visit(v.Field(i), seen)
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			return visit(key, seen) || visit(v.MapIndex(key), seen)
		}
	case reflect.Ptr, reflect.Interface:
		return visit(v.Elem(), seen)
	}
	return false
}
