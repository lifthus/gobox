package tmpEncoder

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Person struct {
	Name string `json:"name"`
	// with Structure tag which is a string written after structure item, we set the rule to handle JSON
	Age int `json:"age"`
} // without tag, there's another specific rule to deal with JSON.

func Enc(toFile Person) (*os.File, func(), error) {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "sample-")
	if err != nil {
		return nil, nil, err
	}

	err = json.NewEncoder(tmpFile).Encode(toFile)
	if err != nil {
		return nil, nil, err
	}

	err = tmpFile.Close()
	if err != nil {
		return nil, nil, err
	}

	return tmpFile, func() { os.Remove(tmpFile.Name()) }, nil
}

func Dec(tmpFile *os.File, close func()) error {
	defer close()
	tmpFile, err := os.Open(tmpFile.Name())
	if err != nil {
		return err
	}
	var fromFile Person
	err = json.NewDecoder(tmpFile).Decode(&fromFile)
	if err != nil {
		return err
	}
	err = tmpFile.Close()
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", fromFile)
	return nil
}

func UsingEncDec() {
	toFile := Person{
		Name: "Jeonghun",
		Age:  25,
	}
	tmpFile, close, err := Enc(toFile)
	if err != nil {
		panic(err)
	}
	err = Dec(tmpFile, close)
	if err != nil {
		panic(err)
	}
}
