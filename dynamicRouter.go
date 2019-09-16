package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// 注册一个动态路由
	// 可以匹配/dynamic/xiaoli
	// 不能匹配/dynamic 和 /dynamic/
	router.GET("/dynamic/:name", testDynamic)

	router.GET("/dynamic1/:name/*action", highDynamic)
	router.Run(":1023")
}

func testDynamic(c *gin.Context)  {
	// 使用 c.param(key) 获取url 参数
	name := c.Param("name")
	c.String(http.StatusOK, "hello, %s", name)
}

func highDynamic(c *gin.Context)  {
	name := c.Param("name")
	action := c.Param("action")
	message := name + " is " + action
	c.String(http.StatusOK, message)
}
