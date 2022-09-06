package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
	"github.com/manaty226/sample-app-with-spicedb/internal/entity"
)

func TestGetBlog(t *testing.T) {
	t.Parallel()
	type want struct {
		blog   entity.Blog
		status int
	}
	tests := map[string]struct {
		path int
		want want
	}{
		"ok": {
			path: 1,
			want: want{
				blog: entity.Blog{
					ID:      1,
					Title:   "test",
					Content: "test",
				},
				status: http.StatusOK,
			},
		},
	}

	for n, tt := range tests {
		tt := tt
		t.Run(n, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Params = []gin.Param{
				{
					Key:   "id",
					Value: fmt.Sprintf("%d", tt.path),
				},
			}
			c.Request, _ = http.NewRequest(
				"GET",
				fmt.Sprintf("/blogs/%d", tt.path),
				bytes.NewReader([]byte("")),
			)
			moq := &GetBlogServiceMock{}
			moq.GetBlogFunc = func(id int) (*entity.Blog, error) {
				return &entity.Blog{
					ID:      1,
					Title:   "test",
					Content: "test",
				}, nil
			}
			sut := GetBlog{Service: moq}
			sut.Handle(c)
			var got, want any
			if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
				t.Fatalf("cannot unmarshal got: %q %v", got, err)
			}
			j, _ := json.Marshal(tt.want.blog)
			if err := json.Unmarshal(j, &want); err != nil {
				t.Fatalf("cannot unmarshal want: %q %v", want, err)
			}
			if diff := cmp.Diff(got, want); diff != "" {
				t.Errorf("got differs: (-got, +want \n %s", diff)
			}
			if tt.want.status != w.Code {
				t.Errorf("want status %d, but got %d", tt.want.status, w.Code)
			}
		})
	}
}
