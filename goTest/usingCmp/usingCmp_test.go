package usingCmp_test

import (
	"github.com/google/go-cmp/cmp"
	"testing"

	"goTest/usingCmp" // when using packagename with _test, only exposed resources can be used by importing.
)

func TestCreatePerson(t *testing.T) {
	expected := usingCmp.Person{
		Name: "Dennis",
		Age:  37,
	}
	result := usingCmp.CreatePerson("Denis", 37)
	if diff := cmp.Diff(expected, result); diff != "" { // reeturns description of difference.
		t.Error(diff) // if diff is empty, means they are equal.
	}
}

func TestCreatePersonIgnoringDate(t *testing.T) {
	comparer := cmp.Comparer(func(x, y usingCmp.Person) bool {
		return x.Name == y.Name && x.Age == y.Age
	})
	expected := usingCmp.Person{
		Name: "Dennis",
		Age:  37,
	}
	result := usingCmp.CreatePerson("Dennis", 37)
	if diff := cmp.Diff(expected, result, comparer); diff != "" { // reeturns description of difference.
		t.Error(diff) // if diff is empty, means they are equal.
	}
}
