package streamEncoder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"handleJSON/tmpEncoder"
)

type Person tmpEncoder.Person

func UsingStreamEncoder() {
	data := `{"name":"Fred", "age":40}
	{"name":"Mary", "age":21}
	{"name":"Pat", "age":30}`

	var t Person

	fmt.Println(" streamEncoder : Decoding ")
	dec := json.NewDecoder(strings.NewReader(data))
	for dec.More() { // saving each single jason at t at each time.
		err := dec.Decode(&t)
		if err != nil {
			panic(err)
		}
		// Processing Decoded json
		fmt.Println(t)
	}

	fmt.Println(" streamEncoder : Encoding ")
	data2 := make([]Person, 0)
	data2 = append(data2, Person{Name: "GOPHER", Age: 13}, Person{Name: "GOLANG", Age: 31})
	// Encoding
	var b bytes.Buffer
	enc := json.NewEncoder((&b))
	for _, input := range data2 {
		t = input
		err := enc.Encode(t)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println(b.String())
}
