package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("/*.html")               // 添加入口index.html
	r.LoadHTMLFiles("/*")                   // 添加资源路径
	r.StaticFile("/hello/", "./index.html") //前端接口

	r.Run(":8080")
}
