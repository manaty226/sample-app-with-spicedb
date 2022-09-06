package store

import (
	"testing"

	"github.com/manaty226/sample-app-with-spicedb/internal/entity"
)

func TestStore(t *testing.T) {
	s := Blogs
	newBlog := &entity.Blog{
		Title:   "test",
		Content: "test",
	}
	s.AddBlog(newBlog)
	if newBlog.ID != 1 {
		t.Errorf("cannot add blog to store")
	}
}
