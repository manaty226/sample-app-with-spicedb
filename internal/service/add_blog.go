package service

import (
	"github.com/manaty226/sample-app-with-spicedb/internal/entity"
)

type AddBlogService struct {
	Store *BlogRepository
}

func (a *AddBlogService) AddBlog(title, content string) (*entity.Blog, error) {
	blog := &entity.Blog{
		Title:   title,
		Content: content,
	}
	err := (*a.Store).AddBlog(blog)
	if err != nil {
		return nil, err
	}
	return blog, nil
}
