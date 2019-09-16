package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func firstGinTest(c *gin.Context) {
	c.String(http.StatusOK, "hello gin")
}

func SecondGinTest(c *gin.Context)  {
	result := map[string]string{"hello":"gin"}
	res := []int{1,2,3,4,5}
	c.String(http.StatusOK, fmt.Sprintf("%v, %v", result, res))
}

func main() {
	engine := gin.Default()

	engine.GET("/helloGin", firstGinTest)

	engine.POST("/postGin", SecondGinTest)

	engine.Run("127.0.0.1:2399")
}
