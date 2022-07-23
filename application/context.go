package application

import (
	"reflect"

	"github.com/Ethan3600/funwithgolang/db"
	"github.com/Ethan3600/funwithgolang/db/adapters"
	"github.com/Ethan3600/funwithgolang/repositories"
)

type Thing interface{}

type ThingFactory func() Thing

type AppContext struct {
	Version    string
	DbStrategy db.DatabaseStrategy
    Container  map[string]Thing
}

func NewApplication() AppContext {
	container := boostrap()

    var db = container["database"].(db.Database)

	return AppContext {
		"1.0.0",
		db.GetStrategy(),
        container,
	}
}

func (ac AppContext) Get(name string) Thing {
    thing := ac.Container[name]  

    if reflect.TypeOf(thing).Kind() == reflect.Func {
        return thing.(ThingFactory)()
    } else {
        return thing
    }
}

func boostrap() map[string]Thing {
    things := make(map[string]Thing)

    inMemAdapter := adapters.NewInMemoryAdapater()
    
	database := db.NewDatabase(inMemAdapter)

	personRepo := repositories.NewPersonRepository(database)

    things["adapter"] = inMemAdapter
    things["database"] = database
    things["personRepository"] = personRepo
    
    return things
}

