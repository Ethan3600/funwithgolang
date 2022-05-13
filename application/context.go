package application

import (
	"echoapp/db"
	"echoapp/db/adapters"
)

type AppContext struct {
	Db         db.Database
	Version    string
	DbStrategy string
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
