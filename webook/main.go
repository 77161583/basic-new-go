package main

import (
	"basic-new-go/webook/internal/web"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	server := gin.Default()

	server.Use(func(ctx *gin.Context) {
		println("第一个middleware")
	})

	server.Use(func(ctx *gin.Context) {
		println("第二个middleware")
	})

	//这里的use 会作用域于当前server的全部路由
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://foo.com"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	u := web.NewUserHandler()
	u.RegisterRoutes(server)
	server.Run()
}
