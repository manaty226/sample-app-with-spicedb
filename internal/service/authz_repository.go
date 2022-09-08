package service

import "github.com/manaty226/sample-app-with-spicedb/internal/entity"

type AuthzRepository interface {
	CheckPermission(object, subject entity.Object, action entity.Action) (bool, error)
	CreateRelation(object, subject entity.Object, relation string) error
}
