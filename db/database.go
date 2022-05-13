package db

type Entity interface {
	GetId() string
}

type Persistence interface {
	Save(e Entity) (string, error)
	GetEntities() ([]Entity, error)

	GetStrategy() string
}

type Database struct {
	adapter Persistence
}

func NewDatabase(p Persistence) Database {
	return Database{p}
}

func (d Database) Save(e Entity) (string, error) {
	return d.adapter.Save(e)
}

func (d Database) GetEntities() ([]Entity, error) {
	return d.adapter.GetEntities()
}

func (d Database) GetStrategy() string {
	return d.adapter.GetStrategy()
}
