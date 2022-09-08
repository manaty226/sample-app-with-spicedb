package store

import (
	"testing"

	"github.com/manaty226/sample-app-with-spicedb/internal/entity"
	"github.com/stretchr/testify/require"
)

func TestStore(t *testing.T) {
	s := Blogs
	newBlog := &entity.Blog{
		Title:   "test",
		Content: "test",
	}
	err := s.AddBlog(newBlog)
	require.NoError(t, err)
	if newBlog.ID != 1 {
		t.Errorf("cannot add blog to store")
	}
}
