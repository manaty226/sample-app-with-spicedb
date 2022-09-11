package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddBlog struct {
	Service    AddBlogService
	Authorizer *Authorizer
}
type PostedBlog struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

func (a *AddBlog) Handle(c *gin.Context) {
	var blog PostedBlog
	if err := c.BindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add blog to store
	b, err := a.Service.AddBlog(blog.Title, blog.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Set authorization user
	id := fmt.Sprintf("%d", b.ID)
	user := c.MustGet(gin.AuthUserKey).(string)
	if err := (*a.Authorizer).CreateUserPermission("blog", id, user, "writer"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := (*a.Authorizer).CreateUserPermission("blog", id, "*", "reader"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": b.ID})
}
