package main

import (
	"basic-new-go/webook/internal/repository"
	"basic-new-go/webook/internal/repository/dao"
	"basic-new-go/webook/internal/service"
	"basic-new-go/webook/internal/web"
	"basic-new-go/webook/internal/web/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	db := initDB()
	server := initWebServer()
	u := initUser(db)
	u.RegisterRoutes(server)
	server.Run(":8080")
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
	//server.Use(cors.New(cors.Config{
	//	//AllowOrigins: []string{"http://172.27.24.126:3000/"},
	//	AllowMethods: []string{"POST", "GET"},
	//	AllowHeaders: []string{"Content-Type", "application/x-www-form-urlencoded;charset=UTF-8"},
	//	//是否允许你带cookie
	//	AllowCredentials: true,
	//	AllowOriginFunc: func(origin string) bool {
	//		if strings.HasPrefix(origin, "http://172.30.64.1:3000/") {
	//			return true
	//		}
	//		return strings.Contains(origin, "yourcompany.com")
	//	},
	//	MaxAge: 12 * time.Hour,
	//}))

	//登录/步骤1
	//store := memstore2.NewStore(
	//	[]byte("%WMN&RKk6rYR1CDcihlEOCjNOGn#GL^G"),
	//	[]byte("LtEnUv74mOe17njL5tClRTsN%ymQTZr7"),
	//)
	store, err := redis.NewStore(
		16, // 最大连接池大小
		"tcp",
		"localhost:6379",
		"", // Redis 密码
		[]byte("2tbTPFVQR4Io8DymwX4H6BECfoswiZ8Y"),
		[]byte("nHNT6o6NQH785UrN35bRIau7toG2YJVM"),
	)
	if err != nil {
		log.Fatalf("Failed to initialize Redis store: %v", err)
	}
	//store := cookie.NewStore([]byte("secret"))
	server.Use(sessions.Sessions("mysession", store))
	//登录/步骤3
	server.Use(middleware.NewLoginMiddlewareBuilder().
		IgnorePaths("/users/login").
		IgnorePaths("/users/signup").
		Build())
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
