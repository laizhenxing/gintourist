package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	// middleware 中间件对 /v1/test1、/v1/test2、/v1/test3 都起作用
	// middleware3 中间件只对当前路由 /v1/test1 起作用
	v1 := router.Group("/v1", middleware)
	{
		v1.GET("/test1", middleware3, test1)
		v1.GET("/test2", test2)
		v1.GET("/test3", test3)
	}
	router.Run(":2201")
}

func middleware(c *gin.Context)  {
	c.String(http.StatusOK, "| exec middleware |")

	c.Next()
}

func test1(c *gin.Context)  {
	c.String(200, c.FullPath())
}

func test2(c *gin.Context)  {
	c.String(200, c.FullPath())
}

func test3(c *gin.Context)  {
	c.String(200, c.FullPath())
}

func middleware3(c *gin.Context)  {
	c.String(200, "| exec middleware3 | ")
	c.Next()
}