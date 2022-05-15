package dtos

import (
	"github.com/Ethan3600/funwithgolang/entities"
)

type Person struct {
	Id    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Age   int    `json:"age,omitempty"`
	Skill string `json:"skill,omitempty"`
}

func (p Person) ToEntity() entities.Person {
	return entities.NewPerson(p.Name, p.Age, &p.Skill)
}
