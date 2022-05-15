package application

import (
	"github.com/Ethan3600/funwithgolang/db"
	"github.com/Ethan3600/funwithgolang/db/adapters"
)

type AppContext struct {
	Db         db.Database
	Version    string
	DbStrategy db.DatabaseStrategy
}

func NewApplication() AppContext {
	database := db.NewDatabase(
		adapters.NewInMemoryAdapater(),
	)

	return AppContext{
		database,
		"1.0.0",
		database.GetStrategy(),
	}
}
