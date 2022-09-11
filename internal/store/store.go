package store

import (
	"fmt"
	"sync"

	"github.com/manaty226/sample-app-with-spicedb/internal/entity"
)

var (
	Blogs = &BlogStore{Blogs: map[int]*entity.Blog{}}
)

type BlogStore struct {
	LastID int
	Blogs  map[int]*entity.Blog
	mu     sync.Mutex
}

func (s *BlogStore) AddBlog(blog *entity.Blog) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.LastID++
	blog.ID = s.LastID
	s.Blogs[s.LastID] = blog
	return nil
}

func (s *BlogStore) GetBlog(id int) (*entity.Blog, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	b, ok := s.Blogs[id]
	if !ok {
		return nil, fmt.Errorf("blog id %d does not exist.", id)
	}
	return b, nil
}

func (s *BlogStore) UpdateBlog(blog *entity.Blog) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.Blogs[blog.ID]; !ok {
		return fmt.Errorf("The requested blog whose id=%d does not exist.", blog.ID)
	}
	s.Blogs[blog.ID] = blog
	return nil
}
