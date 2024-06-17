package main

import (
	"basic-new-go/webook/internal/repository"
	"basic-new-go/webook/internal/repository/dao"
	"basic-new-go/webook/internal/service"
	"basic-new-go/webook/internal/web"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
)

func main() {
	db := initDB()
	server := initWebServer()
	u := initUser(db)
	u.RegisterRoutes(server)
	server.Run()
}

func initWebServer() *gin.Engine {
	server := gin.Default()
	server.Use(func(ctx *gin.Context) {
		println("第一个middleware")
	})
	server.Use(func(ctx *gin.Context) {
		println("第二个middleware")
	})
	//这里的use 会作用域于当前server的全部路由
	server.Use(cors.New(cors.Config{
		//AllowOrigins: []string{"http://172.27.24.126:3000/"},
		AllowMethods: []string{"POST", "GET"},
		AllowHeaders: []string{"Content-Type", "application/x-www-form-urlencoded;charset=UTF-8"},
		//是否允许你带cookie
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://172.27.24.126:3000/") {
				return true
			}
			return strings.Contains(origin, "yourcompany.com")
		},
		MaxAge: 12 * time.Hour,
	}))
	return server
}

func initUser(db *gorm.DB) *web.UserHandler {
	ud := dao.NewUserDao(db)
	repo := repository.NewUserRepository(ud)
	svc := service.NewUserService(repo)
	u := web.NewUserHandler(svc)
	return u
}

func initDB() *gorm.DB {
	dsn := "root:root@tcp(localhost:13316)/webook"
	db, error := gorm.Open(mysql.Open(dsn))
	if error != nil {
		//初始化过成中panic
		panic(error)
	}
	err := dao.InitTable(db)
	if err != nil {
		panic(err)
	}
	return db
}
