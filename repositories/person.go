package repositories

import (
	"echoapp/db"
	"echoapp/entities"
	"sync"
)

var once sync.Once

var singleton *PersonRepository

type PersonRepository struct {
	db db.Database
}

func NewPersonRepository(db db.Database) PersonRepository {

	once.Do(func() {
		singleton = &PersonRepository{
			db,
		}
	})

	return *singleton
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
