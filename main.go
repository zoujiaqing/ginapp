package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zoujiaqing/ginapp/routers"
)

func main() {
	r := gin.Default()
	routers.LoadUserRouters(r)
	r.Run(":8080") // 可以从本地的127.0.0.1:8080访问，不填的话，默认是8080端口
}
