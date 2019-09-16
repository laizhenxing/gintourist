package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// 注册路由和handler
	// url: /get/params/?firstname=xing&lastname=xiaoli
	router.GET("/get/params", func(c *gin.Context) {
		// 获取参数内容，获取的所有参数的类型都是string
		// DefaultQuery(key, default) 获取参数key, 如果不存在，则使用第二个参数作为默认值
		firstname := c.DefaultQuery("firstname", "guest")
		// Query(key) 获取参数key的值，没有默认值
		lastname := c.Query("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	router.Run(":1022")
}
