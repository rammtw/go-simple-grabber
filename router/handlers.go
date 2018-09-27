package router

import (
	"github.com/gin-gonic/gin"
	"lesson/news"
	"net/http"
)

func indexHandler(c *gin.Context)  {
	c.String(http.StatusOK, "Hello, World!")
}

func collectHandler(c *gin.Context)  {
	category := c.Param("category")

	news.Collect(category)
	c.String(http.StatusOK, "Search is initialized!")
}

func resultHandler(c *gin.Context)  {
	category := c.Param("category")

	topics := news.Result(category)
	c.JSON(http.StatusOK, topics)
}
