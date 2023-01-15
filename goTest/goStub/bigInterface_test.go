package goStub

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestLogicGetPetNames(t *testing.T) {
	data := []struct {
		name     string
		userID   string
		petNames []string
	}{
		{"case1", "1", []string{"Bubbles"}},
		{"case2", "2", []string{"Stampy", "Snowall II"}},
		{"case3", "3", nil},
	}
	l := Logic{GetPetNamesStub{}}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			petNames, err := l.GetPetNames(d.userID)
			if err != nil {
				t.Error(err)
			}
			if diff := cmp.Diff(d.petNames, petNames); diff != "" {
				t.Error(diff)
			}
		})
	}
}

// ======

func TestLogicGetPetNames(t *testing.T) {
	data := []struct {
		name     string
		getPets  func(userID string) ([]Pet, error)
		userID   string
		petNames []string
		errMsg   string
	}{
		{"case1", func(userID string) ([]Pet, error) {
			return []Pet{{Name: "Bubbles"}}, nil
		}, "1", []string{"Bubbles"}, ""},
		{"case2", func(userID string) ([]Pet, error) {
			return nil, errors.New("invalid id: 3")
		}, "3", nil, "invalid id: 3"},
	}
	l := Logic{}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			l.Entities = EntitiesStub{getPets: d.getPets}
			petNames, err := l.GetPetNaems(d.userID)
			if diff := cmp.Diff(petNames, d.petNames); diff != "" {
				t.Error(diff)
			}
			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if errMsg != d.errMsg {
				t.Errorf("Expected error `%s`, got `%s`", d.errMsg, errMsg)
			}
		})
	}
}
