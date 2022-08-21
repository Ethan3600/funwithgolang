package person

import (
	"github.com/Ethan3600/funwithgolang/dtos"
	"github.com/Ethan3600/funwithgolang/entities"
	"github.com/Ethan3600/funwithgolang/repositories"
)

func CreatePerson(p dtos.Person, repo repositories.PersonRepository) (string, error) {
	personEntity := p.ToEntity()

	return repo.Save(personEntity)
}

func GetPeople(repo repositories.PersonRepository) ([]dtos.Person, error) {
	peopleEntities, err := repo.GetPeople()
	if err != nil {
		return nil, err
	}

	return from_entities(peopleEntities), nil
}

func GetPerson(id string, repo repositories.PersonRepository) (*dtos.Person, error) {
	var person *dtos.Person

	personEntity, err := repo.GetPerson(id)
	if err != nil {
		return nil, err
	}

	if personEntity != nil {
		dto := from_entity(*personEntity)
		person = &dto
	} else {
		person = nil
	}

	return person, nil
}

func from_entities(peopleEntities []entities.Person) []dtos.Person {
	var people []dtos.Person
	for _, p := range peopleEntities {
		people = append(people, from_entity(p))
	}
	return people
}

func from_entity(p entities.Person) dtos.Person {
	return dtos.Person{
		Id:    p.Id,
		Name:  p.Name,
		Age:   p.Age,
		Skill: p.Skill,
	}
}
