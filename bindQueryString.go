package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

// 定义一个结构体，用来绑定 url query
type Person struct {
	Name string `form:"name"`
	Address string `form:"address"`
}

func main() {
	router := gin.Default()
	router.Any("/testing", startPage)
	router.Run(":1055")
}

func startPage(c *gin.Context)  {
	var person Person
	// 将 url 查询参数和 person 绑定在一起
	if c.ShouldBindQuery(&person) == nil {
		log.Println("==== Only Bind by Query String ====")
		log.Println(c.Param("name"))
		log.Println("person: ", person)
		log.Println(person.Name)
		log.Println(person.Address)
	}
	c.String(200, "Success")
}
