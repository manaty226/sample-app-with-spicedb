package handler

import (
	"github.com/manaty226/sample-app-with-spicedb/internal/entity"
	"github.com/manaty226/sample-app-with-spicedb/internal/service"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . AddBlogService GetBlogService CheckPermissionService CreateRelation
type AddBlogService interface {
	AddBlog(title, content string) (*entity.Blog, error)
}

type GetBlogService interface {
	GetBlog(id int) (*entity.Blog, error)
}

type Authorizer interface {
	CheckPermission(objectType, objectId, user string, method service.Method) (bool, error)
	CreateUserPermission(objectType, objectId, user, relation string) error
}
