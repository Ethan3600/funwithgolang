package dtos

import "echoapp/db"

type Dto interface {
	ToEntity() db.Entity
}
