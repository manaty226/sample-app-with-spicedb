package handler

import (
	"github.com/manaty226/sample-app-with-spicedb/internal/entity"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . AddBlogService
type AddBlogService interface {
	AddBlog(title, content string) (entity.Blog, error)
}
