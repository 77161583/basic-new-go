package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()
	//路由
	server.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})

	server.POST("/post", func(c *gin.Context) {
		c.String(http.StatusOK, "I am post")
	})

	server.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "这是参数路由"+name)
	})

	server.GET("/views/*.html", func(c *gin.Context) {
		path := c.Param(".html")
		c.String(http.StatusOK, "这是通配符路由 %s"+path)
	})

	server.GET("/order", func(c *gin.Context) {
		oid := c.Query("id")
		c.String(http.StatusOK, "当前id："+oid)
	})

	//端口从配置文件里读
	server.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
