package service

import (
	"github.com/manaty226/sample-app-with-spicedb/internal/entity"
)

type Authorizer struct {
	Store *AuthzRepository
}

type Method string

const (
	GET  Method = "GET"
	POST Method = "POST"
	PUT  Method = "PUT"
)

func (m Method) Method2Action() entity.Action {
	switch m {
	case GET:
		return entity.Read
	case POST:
	case PUT:
		return entity.Write
	}
	return entity.Read
}

func (a Authorizer) CheckPermission(
	objectType string,
	objectID string,
	user string,
	method Method,
) (bool, error) {
	object := entity.Object{Type: entity.ObjectType(objectType), ID: objectID}
	subject := entity.Object{Type: entity.ObjectType("user"), ID: user}
	action := method.Method2Action()

	ok, err := (*a.Store).CheckPermission(object, subject, action)
	if err != nil {
		return false, err
	}
	return ok, nil
}

func (a *Authorizer) CreateUserPermission(
	objectType string,
	objectId string,
	user string,
	relation string,
) error {
	object := entity.Object{Type: entity.ObjectType(objectType), ID: objectId}
	subject := entity.Object{Type: entity.ObjectType("user"), ID: user}

	if err := (*a.Store).CreateRelation(object, subject, relation); err != nil {
		return err
	}
	return nil
}
