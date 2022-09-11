package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateBlog struct {
	Service    UpdateBlogService
	Authorizer *Authorizer
}

type UpdatedBlog struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

func (u *UpdateBlog) Handle(c *gin.Context) {
	// Check authorization to get blog
	user := c.MustGet(gin.AuthUserKey).(string)
	ok, err := (*u.Authorizer).CheckPermission("blog", c.Param("id"), user, "PUT")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !ok {
		c.JSON(http.StatusForbidden, gin.H{"error": fmt.Sprintf("not authorized to update the requested blog.")})
		return
	}

	var blog UpdatedBlog
	if err := c.BindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Update blog to store
	b, err := u.Service.UpdateBlog(id, blog.Title, blog.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, b)
}
