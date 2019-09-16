package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// gin.H 本质是 map[string]interface{}
	router.GET("/test/json", func(c *gin.Context) {
		// 会输出头格式为 application/json; charset=UTF-8 的 json 字符串
		c.JSON(http.StatusOK, gin.H{"msg":"test json format", "code": http.StatusOK})
	})

	router.GET("/test/moreJson", func(c *gin.Context) {
		// 直接使用结构体定义
		var msg struct{
			Name string `json:"user"`
			Message string
			Number int
		}
		msg.Name = "john"
		msg.Message = "test more json"
		msg.Number = 122
		// 输出结果
		// {"user": "john", "Message": "test more json","Number": 122}
		c.JSON(http.StatusOK, msg)
	})
	
	router.GET("/test/xml", func(c *gin.Context) {
		// 输出头格式(Content-Type):application/xml; charset=utf-8
		// 输出结果是 xml 字符串
		// <map>
		//    <msg>test xml format</msg>
		//    <code>200</code>
		//</map>
		c.XML(http.StatusOK, gin.H{"msg":"test xml format", "code":http.StatusOK})
	})

	router.GET("/test/yaml", func(c *gin.Context) {
		// 输出头格式(Content-Type): application/x-yaml; charset=utf-8
		// 输出结果
		// code: 200
		//msg: test yaml format
		c.YAML(http.StatusOK, gin.H{"msg":"test yaml format", "code":http.StatusOK})
	})
	router.Run(":1244")
}
