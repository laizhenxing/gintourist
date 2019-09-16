package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.POST("/test/post", func(c *gin.Context) {
		// 获取post传的message的内容
		// c.PostParam获取的所有参数内容的类型都但是string
		message := c.PostForm("message")
		// 使用默认值
		nick := c.DefaultPostForm("nick", "default_nick")
		// 使用json返回
		c.JSON(200, gin.H{
			"status": "posted",
			"message": message,
			"nick": nick,
		})
	})
	router.Run(":1014")
}
