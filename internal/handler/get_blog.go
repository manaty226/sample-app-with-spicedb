package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetBlog struct {
	Service GetBlogService
}

func (g GetBlog) Handle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	b, err := g.Service.GetBlog(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, b)
}
