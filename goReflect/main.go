package main

import (
	"encoding/csv"
	"fmt"
	"goReflect/basicReflect"
	"goReflect/marshallingCSV"
	"strings"
)

func main() {
	basicReflect.UseBasicReflect()
	basicReflect.UseNewValue()
	basicReflect.CheckNil()

	fmt.Println("=== Marshalling ===")
	data := `name,age,has_pet
Jon,"100",true
"Fred ""The Hammer"" Smith", 42, false
Martha,37,"true"
`

	r := csv.NewReader(strings.NewReader(data))
	allData, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	var entries []marshallingCSV.MyData
	marshallingCSV.Unmarshal(allData, &entries)
	fmt.Println(entries)

	out, err := marshallingCSV.Marshal(entries)
	if err != nil {
		panic(err)
	}
	sb := &strings.Builder{}
	w := csv.NewWriter(sb)
	w.WriteAll(out)
	fmt.Println(sb)
}
