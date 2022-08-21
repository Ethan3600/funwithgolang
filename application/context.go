package application

import (
	"github.com/Ethan3600/funwithgolang/db"
	"github.com/Ethan3600/funwithgolang/db/adapters"
	"github.com/Ethan3600/funwithgolang/repositories"
)

type Container struct {
	Adapter    db.Persistence
	Database   db.Database
	PersonRepo repositories.PersonRepository
}

type AppContext struct {
	Version    string
	DbStrategy db.DatabaseStrategy
	C          Container
}

func NewApplication() AppContext {
	container := boostrap()

	var db = container.Database

	return AppContext{
		"1.0.0",
		db.GetStrategy(),
		container,
	}
}

func boostrap() Container {
	c := Container{}

	inMemAdapter := adapters.NewInMemoryAdapater()

	database := db.NewDatabase(inMemAdapter)

	personRepo := repositories.NewPersonRepository(database)

	c.Adapter = inMemAdapter
	c.Database = database
	c.PersonRepo = personRepo

	return c
}
