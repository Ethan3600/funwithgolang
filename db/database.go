package db

type Entity interface {
	GetId() string
}

type Persistence interface {
	Save(e Entity) (string, error)
	GetEntities() ([]Entity, error)
	GetEntity(id string) (*Entity, error)

	GetStrategy() DatabaseStrategy
}

type DatabaseStrategy string

const (
	InMemory DatabaseStrategy = "in_memory"
)

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

func (d Database) GetEntity(id string) (*Entity, error) {
	return d.adapter.GetEntity(id)
}

func (d Database) GetStrategy() DatabaseStrategy {
	return d.adapter.GetStrategy()
}
