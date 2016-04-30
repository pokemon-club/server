package controller

type CRUDController interface {
	All() interface{}
	Find(id string) (interface{}, error)
}
