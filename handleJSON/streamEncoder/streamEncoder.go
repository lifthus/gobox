package streamEncoder

import (
	"encoding/json"
	"fmt"
	"strings"

	"handleJSON/tmpEncoder"
)

func UsingStreamEncoder() {
	data := `{"name":"Fred", "age":40}
	{"name":"Mary", "age":21}
	{"name":"Pat", "age":30}`

	var t tmpEncoder.Person

	dec := json.NewDecoder(strings.NewReader(data))
	for dec.More() { // saving each single jason at t at each time.
		err := dec.Decode(&t)
		if err != nil {
			panic(err)
		}
		// Processing Decoded json
		fmt.Println(t)
	}
}
