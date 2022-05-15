package dtos

import "github.com/Ethan3600/funwithgolang/db"

type Dto interface {
	ToEntity() db.Entity
}
