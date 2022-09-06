package service

import (
	"github.com/manaty226/sample-app-with-spicedb/internal/entity"
	"github.com/manaty226/sample-app-with-spicedb/internal/handler"
)

var _ handler.GetBlogService = (*GetBlogService)(nil)

type GetBlogService struct {
	Store *BlogRepository
}

func (g *GetBlogService) GetBlog(id int) (*entity.Blog, error) {
	b, err := (*g.Store).GetBlog(id)
	if err != nil {
		return nil, err
	}
	return b, nil
}
