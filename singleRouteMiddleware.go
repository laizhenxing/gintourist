package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()

	router.GET("/singleRoute/middleware", middleware1, middleware2, handler)

	router.Run(":2322")
}

func middleware1(c *gin.Context)  {
	log.Println("exec middleware1")

	// 执行下一步操作
	c.Next()
}

func middleware2(c *gin.Context)  {
	log.Println("arrive at middleware2")
	// 执行下一步操作
	c.Next()

	// 再执行 middleware2
	log.Println("exec middleware2")
}

func handler(c *gin.Context)  {
	log.Println("exec handler")
}

