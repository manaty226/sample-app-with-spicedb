package store

import (
	"fmt"

	"github.com/manaty226/sample-app-with-spicedb/internal/entity"
	"github.com/manaty226/sample-app-with-spicedb/internal/service"
)

var (
	_     service.BlogRepository = (*BlogStore)(nil)
	Blogs                        = &BlogStore{Blogs: map[int64]*entity.Blog{}}
)

type BlogStore struct {
	LastID int64
	Blogs  map[int64]*entity.Blog
}

func (s *BlogStore) AddBlog(blog *entity.Blog) error {
	s.LastID++
	blog.ID = s.LastID
	s.Blogs[s.LastID] = blog
	return nil
}

func (s *BlogStore) GetBlog(id int) (*entity.Blog, error) {
	b, ok := s.Blogs[int64(id)]
	if !ok {
		return nil, fmt.Errorf("blog id %d does not exist.", id)
	}
	return b, nil
}
