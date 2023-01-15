package goStub

import "fmt"

type Person int
type User Person
type Pet struct {
	Name string
}

type Entities interface {
	GetUser(id string) (User, error)
	GetPets(userID string) ([]Pet, error)
	GetChildren(userID string) ([]Person, error)
	GetFriends(UserID string) ([]Person, error)
	SaveUser(user User) error
}

type Logic struct {
	Entities Entities
}

// assume the method below should be tested.
func (l Logic) GetPetNames(userId string) ([]string, error) {
	pets, err := l.Entities.GetPets(userId)
	if err != nil {
		return nil, err
	}
	out := make([]string, len(pets))
	for _, p := range pets {
		out = append(out, p.Name)
	}
	return out, nil
}

type GetPetNamesStub struct {
	Entities
}

func (ps GetPetNamesStub) GetPets(userID string) ([]Pet, error) {
	switch userID {
	case "1":
		return []Pet{{Name: "Bubbles"}}, nil
	case "2":
		return []Pet{{Name: "Stamply"}, {Name: "Snowball II"}}, nil
	default:
		return nil, fmt.Errorf("invalid id: %s", userID)
	}
}

// ======

type EntitiesStub struct {
	getUser     func(id string) (User, error)
	getPets     func(userID string) ([]Pet, error)
	getChildren func(userID string) ([]Pet, error)
	getFriends  func(userID string) ([]Person, error)
	saveUser    func(user User) error
}

func (es EntitiesStub) GetUser(id string) (User, error) {
	return es.getUser(id)
}

func (es EntitiesStub) GetPets(userID string) ([]Pet, error) {
	return es.getPets(userID)
}
