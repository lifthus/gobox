package basicReflect

import (
	"fmt"
	"reflect"
)

func hasNoValue(i interface{}) bool {
	iv := reflect.ValueOf(i)
	if !iv.IsValid() {
		return true
	}
	switch iv.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Func, reflect.Interface:
		return iv.IsNil()
	default:
		return false
	}
}

func CheckNil() {
	fmt.Println("=== Chcke Nil ===")
	type Tmp interface{}
	fmt.Println(hasNoValue(5))

	var a Tmp
	fmt.Println(hasNoValue(a))

	fmt.Println(hasNoValue(nil))

	var b []byte
	fmt.Println(hasNoValue(b))

	var c map[string]int
	fmt.Println(hasNoValue(c))
}
