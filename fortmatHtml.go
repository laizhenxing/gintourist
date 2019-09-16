package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/test/html", testHtml)
	// 注册一个静态的资源目录
	// 该目录下的资源看可以按照路径访问
	router.Static("/static", "./static")
	// 注册一个路径，gin 加载模板的时候从该目录查找
	// 数是一个匹配字符，如 views/*/* 的意思是 模板目录有两层
	router.LoadHTMLGlob("views/*/*")
	router.Run(":1245")
}

func testHtml(c *gin.Context)  {
	// 输出 HTML
	// 第一个参数是状态码
	// 第二个参数("shop/search")是要加载的模板路径,对应目录路径 views/shop/search.html
	// 第三个参数是模板变量
	c.HTML(http.StatusOK, "shop/search", gin.H{"keywords":"test html"})
}
