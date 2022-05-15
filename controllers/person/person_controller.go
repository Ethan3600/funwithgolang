package person

import (
	"github.com/Ethan3600/funwithgolang/application"
	"github.com/Ethan3600/funwithgolang/dtos"
	"github.com/Ethan3600/funwithgolang/entities"
	"github.com/Ethan3600/funwithgolang/repositories"
)

func CreatePerson(p dtos.Person, app application.AppContext) (string, error) {
	personEntity := p.ToEntity()

	repo := repositories.NewPersonRepository(app.Db)

	return repo.Save(personEntity)
}

func GetPeople(app application.AppContext) ([]dtos.Person, error) {
	repo := repositories.NewPersonRepository(app.Db)

	peopleEntities, err := repo.GetPeople()
	if err != nil {
		return nil, err
	}

	return from_entities(peopleEntities), nil
}

func GetPerson(id string, app application.AppContext) (*dtos.Person, error) {
	repo := repositories.NewPersonRepository(app.Db)
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
