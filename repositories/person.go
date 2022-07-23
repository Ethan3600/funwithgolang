package repositories

import (
	"sync"

	"github.com/Ethan3600/funwithgolang/db"
	"github.com/Ethan3600/funwithgolang/entities"
)

var once sync.Once

type PersonRepository struct {
	db db.Database
}

func NewPersonRepository(db db.Database) PersonRepository {

    repo := PersonRepository{
        db,
    }

	return repo
}

func (r PersonRepository) Save(p entities.Person) (string, error) {
	return r.db.Save(p)
}

func (r PersonRepository) GetPeople() ([]entities.Person, error) {
	results, err := r.db.GetEntities()
	if err != nil {
		return nil, err
	}

	return from_entities(results), nil
}

func (r PersonRepository) GetPerson(id string) (*entities.Person, error) {
	result, err := r.db.GetEntity(id)
	if err != nil {
		return nil, err
	}

	if *result != nil {
		person := from_entity(*result)
		return &person, nil
	} else {
		return nil, nil
	}
}

func from_entities(peopleEntities []db.Entity) []entities.Person {
	var people []entities.Person
	for _, p := range peopleEntities {
		people = append(people, from_entity(p))
	}
	return people
}

func from_entity(e db.Entity) entities.Person {
	return e.(entities.Person)
}
