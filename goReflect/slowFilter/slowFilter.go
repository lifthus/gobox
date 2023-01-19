package slowFilter

import (
	"fmt"
	"reflect"
)

func SlowFilter(slice interface{}, filter interface{}) interface{} {
	sv := reflect.ValueOf(slice)
	fv := reflect.ValueOf(filter)

	sliceLen := sv.Len()
	out := reflect.MakeSlice(sv.Type(), 0, sliceLen)
	for i := 0; i < sliceLen; i++ {
		curVal := sv.Index(i)
		values := fv.Call([]reflect.Value{curVal})
		if values[0].Bool() {
			out = reflect.Append(out, curVal)
		}
	}
	return out.Interface()
}

func UsingSlowFilter() {
	fmt.Println("=== Using Slow Filter ===")
	s1 := []string{"Jack", "Jon", "Rain", "Andrew"}
	fmt.Println(s1)
	s2 := SlowFilter(s1, func(s string) bool {
		return len(s) > 3
	}).([]string)
	fmt.Println(s2)
}
