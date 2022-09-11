package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/manaty226/sample-app-with-spicedb/internal/config"
	"github.com/manaty226/sample-app-with-spicedb/internal/handler"
	"github.com/manaty226/sample-app-with-spicedb/internal/service"
	"github.com/manaty226/sample-app-with-spicedb/internal/store"
)

func runner() (*gin.Engine, error) {
	r := gin.Default()
	blogs := service.BlogRepository(store.Blogs)
	cfg, err := config.New()
	if err != nil {
		return nil, err
	}
	spice, err := store.NewSpiceDB(cfg.SpiceHost, cfg.SpicePort)
	if err != nil {
		return nil, err
	}
	if err := spice.CreateNameSpace(cfg.NameSpace); err != nil {
		return nil, err
	}
	authzRepository := service.AuthzRepository(spice)
	authorizer := handler.Authorizer(&service.Authorizer{Store: &authzRepository})

	a := handler.AddBlog{
		Service:    &service.AddBlogService{Store: &blogs},
		Authorizer: &authorizer,
	}
	g := handler.GetBlog{
		Service:    &service.GetBlogService{Store: &blogs},
		Authorizer: &authorizer,
	}
	u := handler.UpdateBlog{
		Service:    &service.UpdateBlogService{Store: &blogs},
		Authorizer: &authorizer,
	}

	authorized := r.Group("/", gin.BasicAuth(cfg.AuthnUsers))
	authorized.POST("/blogs", a.Handle)
	authorized.GET("/blogs/:id", g.Handle)
	authorized.PUT("/blogs/:id", u.Handle)

	return r, nil
}

func main() {
	r, err := runner()
	if err != nil {
		fmt.Printf("Error in setup server: %v", err.Error())
		return
	}
	if err := r.Run(":3000"); err != nil {
		fmt.Printf("Error in run server: %v", err.Error())
		return
	}
}
