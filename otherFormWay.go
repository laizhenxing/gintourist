package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()

	router.POST("/post/other", func(c *gin.Context) {
		// 获取原始字节
		d, err := c.GetRawData()
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(string(d))
		c.String(200, "ok")
	})
	router.Run(":1654")
}
