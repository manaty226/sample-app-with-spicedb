package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddBlog struct {
	Service AddBlogService
}
type PostedBlog struct {
	Title   string `json:"id" validate:"required"`
	Content string `json:"content" validate:"required"`
}

func (a *AddBlog) Handle(c *gin.Context) {
	var blog PostedBlog
	if err := c.BindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	b, err := a.Service.AddBlog(blog.Title, blog.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"id": b.ID})
}
