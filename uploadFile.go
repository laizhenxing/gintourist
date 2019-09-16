package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
)

func main() {
	router := gin.Default()

	// 设置文件上传大小, route.MaxMultipartMemory = 8 << 20
	// 处理单一文件上传
	// 文件参数查看
	router.POST("/upload/params", uploadParam)
	// 文件上传
	router.POST("/upload", uploadFile)

	router.POST("/multi/params", multiUploadParams)

	// 多文件上传
	router.POST("/multi/upload", multiUpload)

	router.Run(":1200")
}

// 单文件上传
func uploadFile(c *gin.Context)  {
	// 利用 c.Request.FormFile 解析文件
	file, header, err := c.Request.FormFile("file")
	log.Println(reflect.TypeOf(file))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "文件上传失败"})
		return
	}
	// 保存文件
	// 创建上传目录
	os.Mkdir("./uploads", os.ModePerm)
	// 创建上传文件
	cur, err := os.Create("./uploads/" + header.Filename)
	defer cur.Close()
	if err != nil {
		log.Fatal(err)
	}

	// 读取文件内容
	content, err := ioutil.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "文件读取失败"})
		return
	}

	// 将读取的文件内容写入新建的文件中
	err = ioutil.WriteFile(cur.Name(), content, os.ModePerm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "文件写入失败"})
		return
	}
	// 把文件数据拷贝到新建的文件
	//io.Copy(cur,file)

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", cur.Name()))
}

// 单文件参数查看
func uploadParam(c *gin.Context)  {
	file, _ := c.FormFile("file")
	log.Println(file.Filename)
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

// 多文件参数查看
func multiUploadParams(c *gin.Context)  {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"msg": "上传文件有误"})
	}
	// 拿到文件集合
	files := form.File["uploads"]
	for _, file := range files {
		log.Println(file.Filename)
	}
	c.JSON(http.StatusOK,
		gin.H{
			"msg": "success",
			"file_len": len(files),
		})
}

// 多文件上传
func multiUpload(c *gin.Context)  {
	// 创建上传目录
	os.Mkdir("./uploads", os.ModePerm)
	// 利用 c.Request.MultipartForm 解析
	// 设置文件大小
	err := c.Request.ParseMultipartForm( 4 << 20)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "文件太大"})
		return
	}

	formdata := c.Request.MultipartForm
	files := formdata.File["files"]

	for _, v := range files {
		file, err := v.Open()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "文件读取失败"})
			return
		}
		defer file.Close()

		content, err := ioutil.ReadAll(file)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "文件读取失败"})
			return
		}

		// 创建文件
		cur, err := os.Create("./uploads/" + v.Filename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "文件创建失败"})
			log.Fatal(err)
			return
		}
		defer cur.Close()

		// 将文件内容写入刚创建的文件中
		err = ioutil.WriteFile(cur.Name(), content, os.ModePerm)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "文件写入失败"})
			log.Fatal(err)
			return
		}
	}
	c.JSON(http.StatusOK , gin.H{"msg": "文件上传成功"})
}