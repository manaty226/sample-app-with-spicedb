package service

import "github.com/manaty226/sample-app-with-spicedb/internal/entity"

type AddBlogService struct {
	Store map[string]*entity.Blog
}

func AddBlog(title, content string) (entity.Blog, error) {
	return entity.Blog{}, nil
}
