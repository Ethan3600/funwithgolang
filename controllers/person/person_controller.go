package person

import (
	"echoapp/application"
	"echoapp/dtos"
	"echoapp/entities"
	"echoapp/repositories"
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

func from_entities(peopleEntities []entities.Person) []dtos.Person {
	var people []dtos.Person
	for _, p := range peopleEntities {
		people = append(people, dtos.Person{
			Id:    p.Id,
			Name:  p.Name,
			Age:   p.Age,
			Skill: p.Skill,
		})
	}
	return people
}
