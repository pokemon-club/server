package controller

import (
	"github.com/pokemon-club/server/model"
)

type CRUDController interface {
	All() interface{}
	Find(id string, m model.Document) (model.Document, error)
	Insert(m model.Document) (model.Document, error)
	Update(id string, m model.Document) (model.Document, error)
	Delete(id string) error
}
