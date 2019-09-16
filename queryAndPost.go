package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// 定义一个 Person 结构体，用来绑定数据
type Person1 struct {
	Name string `form:"name"`
	Address string `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2019-01-01" time_utc:"1"`
}

func main() {
	router := gin.Default()
	router.GET("/testing/query", bindQ)
	router.Run(":1311")
}

func bindQ(c *gin.Context)  {
	var person Person1
	if c.ShouldBind(&person) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
		"data": person,
	})
}
