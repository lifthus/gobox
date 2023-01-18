package main

import (
	"fmt"
	"reflect"
)

func main() {
	type Foo struct {
		A int    `myTag:"value"`
		B string `myTag:"value2"`
	}

	fmt.Println("=== reflection of name ===")
	var x int
	xt := reflect.TypeOf(x)
	fmt.Println(xt.Name()) // int
	f := Foo{}
	ft := reflect.TypeOf(f)
	fmt.Println(ft.Name()) // Foo
	xpt := reflect.TypeOf(&x)
	fmt.Println(xpt.Name()) // empty string. types like sllice and pointer don't have name.
	fmt.Println("=== reflection of reference ===")
	x = 0
	xpt = reflect.TypeOf(&x)
	fmt.Println(xpt.Name())        // empty string
	fmt.Println(xpt.Kind())        // reflect.Ptr
	fmt.Println(xpt.Elem().Name()) // "int"
	fmt.Println(xpt.Elem().Kind()) // reflect.Int
	fmt.Println("=== reflection of struct ===")
	f = Foo{}
	ft = reflect.TypeOf(f)
	for i := 0; i < ft.NumField(); i++ {
		curField := ft.Field(i)
		fmt.Println(curField.Name, curField.Type.Name(), curField.Tag.Get("myTag"))
	}
	fmt.Println("=== value of reflection ===")
	var v int = 3
	vValue := reflect.ValueOf(v)       // reflect.value instance
	fmt.Println(vValue, vValue.Type()) // 3 int
	s := []string{"a", "b", "c"}
	sv := reflect.ValueOf(s)        // sv is of type reflect.Value
	s2 := sv.Interface().([]string) // s2 is of type []string( no type info with Interface so type assertion needed)
	// of course methods like Int, String for primitive types exist.
	fmt.Println(sv, s2)
	fmt.Println("=== Setting value with reflection ===")
	i := 10
	iv := reflect.ValueOf(&i) // reflect.Value instance that represents pointer
	ivv := iv.Elem()          // like Type.Elem(), Value.Elem() returns pointer's indexing value or interface's value.
	ivv.SetInt(20)            // there are methods to set primitive types' value.
	fmt.Println(i)            // 20
}
