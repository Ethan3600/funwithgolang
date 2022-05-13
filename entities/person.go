package entities

import "github.com/google/uuid"

type Person struct {
	Id    string
	Name  string
	Age   int
	Skill string
}

func NewPerson(name string, age int, skill *string) Person {
	default_skill := "talking"

	if skill == nil {
		skill = &default_skill
	}

	id := uuid.New().String()

	return Person{
		id,
		name,
		age,
		*skill,
	}
}

func (p Person) GetId() string {
	return p.Id
}
