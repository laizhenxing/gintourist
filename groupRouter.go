package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// 定义一个路由组(使用花括号，推荐)
	// 定义一个路由前缀 /v1/login 会匹配到这个值
	v1 := router.Group("/v1")
	{
		v1.POST("/login", loginTest)
		v1.POST("/submit", submitTest)
		v1.POST("/read", readTest)
	}

	// 定义一个路由组前缀
	// 不适用花括号
	v2 := router.Group("/v2")
	v2.POST("/login", loginTest)
	v2.POST("/submit", submitTest)
	v2.POST("/read", readTest)

	router.Run(":1011")
}

func loginTest(c *gin.Context)  {
	c.String(http.StatusOK, c.FullPath(), c.ClientIP())
}

func submitTest(c *gin.Context)  {
	c.String(http.StatusOK, c.FullPath())
}

func readTest(c *gin.Context)  {
	c.String(http.StatusOK, c.FullPath())
}
