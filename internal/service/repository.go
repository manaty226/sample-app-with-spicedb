package service

import "github.com/manaty226/sample-app-with-spicedb/internal/entity"

type BlogRepository interface {
	AddBlog(blog *entity.Blog) error
	GetBlog(id int) (*entity.Blog, error)
}
