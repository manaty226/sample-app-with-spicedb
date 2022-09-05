package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
	"github.com/manaty226/sample-app-with-spicedb/internal/entity"
	"github.com/manaty226/sample-app-with-spicedb/testutil"
)

func TestAddBlog(t *testing.T) {
	t.Parallel()

	type want struct {
		status   int
		respFile string
	}
	tests := map[string]struct {
		reqFile string
		want    want
	}{
		"ok": {
			reqFile: "testdata/add_blog/ok_add_blog_req.json",
			want: want{
				status:   http.StatusOK,
				respFile: "testdata/add_blog/ok_add_blog_resp.json",
			},
		},
	}

	for n, tt := range tests {
		tt := tt
		t.Run(n, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request, _ = http.NewRequest(
				"POST",
				"/blogs",
				bytes.NewReader(testutil.LoadFile(t, tt.reqFile)),
			)
			moq := &AddBlogServiceMock{}
			moq.AddBlogFunc = func(title, content string) (entity.Blog, error) {
				return entity.Blog{
					ID:      1,
					Title:   "test",
					Content: "test",
				}, nil
			}
			sut := AddBlog{Service: moq}
			sut.Handle(c)
			var got, want any
			if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
				t.Fatalf("cannot unmarshal got: %q %v", got, err)
			}
			if err := json.Unmarshal(testutil.LoadFile(t, tt.want.respFile), &want); err != nil {
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
