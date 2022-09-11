package service

import (
	"github.com/manaty226/sample-app-with-spicedb/internal/entity"
)

type UpdateBlogService struct {
	Store *BlogRepository
}

func (u *UpdateBlogService) UpdateBlog(id int, title, content string) (*entity.Blog, error) {
	blog := &entity.Blog{
		ID:      id,
		Title:   title,
		Content: content,
	}
	err := (*u.Store).UpdateBlog(blog)
	if err != nil {
		return nil, err
	}
	return blog, nil
}
