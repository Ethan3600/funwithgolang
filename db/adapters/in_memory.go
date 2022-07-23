package adapters

import (
	"github.com/Ethan3600/funwithgolang/db"
)

type InMemory struct {
	data map[string]db.Entity
}

func NewInMemoryAdapater() InMemory {
	data := make(map[string]db.Entity)
	return InMemory{
		data,
	}
}

func (im InMemory) Save(e db.Entity) (string, error) {
	var err error = nil

	uid := e.GetId()

	im.data[uid] = e

	return uid, err
}

func (im InMemory) GetEntities() ([]db.Entity, error) {
	var results []db.Entity

	for _, e := range im.data {
		results = append(results, e)
	}

	return results, nil
}

func (im InMemory) GetEntity(id string) (*db.Entity, error) {
	var result  = im.data[id]

	return &result, nil
}

func (im InMemory) GetStrategy() db.DatabaseStrategy {
	return db.InMemory
}
