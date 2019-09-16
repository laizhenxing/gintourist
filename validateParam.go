package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 定义一个结构体，并指定要验证的数据字段
type Student struct {
	// 参数中必须包含 name
	Name string `form:"name" json:"name" binding:"required"`
	// 参数中必须包含 age
	Age int `form:"age" json:"age" binding:"required"`
}

func main() {
	router := gin.Default()
	router.POST("/testing/validate", validateStu)
	router.Run(":1099")
}

func validateStu(c *gin.Context)  {
	var student Student
	// 将输入的参数绑定到 Student (struct)
	err := c.ShouldBind(&student)
	if err == nil{
		// 绑定成功，进行数据验证
		if student.Name == "xing" && student.Age == 23 {
			c.JSON(http.StatusOK, gin.H{"msg": "validate success"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "validate fail"})
		}
	} else {
		// 绑定失败
		c.JSON(http.StatusBadRequest, gin.H{"msg": "bad request"})
	}
}
