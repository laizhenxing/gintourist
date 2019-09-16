package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getting(c *gin.Context)  {
	c.String(http.StatusOK, "RESTful Api getting")
}
func posting(c *gin.Context)  {
	c.String(http.StatusOK, "RESTful Api posting")
}
func putting(c *gin.Context)  {
	c.String(http.StatusOK, "RESTful Api putting")
}
func deleting(c *gin.Context)  {
	c.String(http.StatusOK, "RESTful Api deleting")
}
func patching(c *gin.Context)  {
	c.String(http.StatusOK, "RESTful Api patching")
}
func heading(c *gin.Context)  {
	c.String(http.StatusOK, "RESTful Api heading")
}
func options(c *gin.Context)  {
	c.String(http.StatusOK, "RESTful Api options")
}

func main() {
	// 初始化引擎
	router := gin.Default()

	router.GET("/getting", getting)
	router.POST("/posting", posting)
	router.PUT("/putting", putting)
	router.DELETE("/deleting", deleting)
	router.PATCH("/patching", patching)
	router.HEAD("/heading", heading)
	router.OPTIONS("/options", options)

	// 默认绑定 :8080
	router.Run()
}
