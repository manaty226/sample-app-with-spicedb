package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetBlog struct {
	Service    GetBlogService
	Authorizer *Authorizer
}

func (g GetBlog) Handle(c *gin.Context) {
	// Check authorization to get blog
	user := c.MustGet(gin.AuthUserKey).(string)
	ok, err := (*g.Authorizer).CheckPermission("blog", c.Param("id"), user, "GET")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !ok {
		c.JSON(http.StatusForbidden, gin.H{"error": fmt.Sprintf("not authorized to get the requested blog.")})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get blog from store
	b, err := g.Service.GetBlog(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, b)
}
