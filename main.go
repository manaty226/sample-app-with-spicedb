package main

import (
	"github.com/gin-gonic/gin"
	"github.com/manaty226/sample-app-with-spicedb/internal/handler"
	"github.com/manaty226/sample-app-with-spicedb/internal/service"
	"github.com/manaty226/sample-app-with-spicedb/internal/store"
)

func runner() *gin.Engine {
	r := gin.Default()
	blogs := service.BlogRepository(store.Blogs)
	a := handler.AddBlog{Service: &service.AddBlogService{Store: &blogs}}
	g := handler.GetBlog{Service: &service.GetBlogService{Store: &blogs}}

	r.POST("/blogs", a.Handle)
	r.GET("/blogs/:id", g.Handle)

	return r
}

func main() {
	r := runner()
	r.Run(":3000")
}
