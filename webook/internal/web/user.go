package web

import (
	"fmt"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserHandler 定义所有跟user用户有关的路由
type UserHandler struct {
	emailExp    *regexp.Regexp
	passwordExp *regexp.Regexp
}

func NewUserHandler() *UserHandler {
	const (
		emailRegex    = "^[a-z0-9_\\.\\-\\+]+@[a-z0-9\\-]+(\\.[a-z0-9\\-]+)*$"
		passwordRegex = `^(?=.*[A-Z])(?=.*[a-z])(?=.*[0-9])(?=.*[!@#$%^&*()_+=-])[A-Za-z0-9!@#$%^&*()_+=-]{8,}$`
	)
	emailExp := regexp.MustCompile(emailRegex, regexp.None)
	passwordExp := regexp.MustCompile(passwordRegex, regexp.None)
	return &UserHandler{
		emailExp:    emailExp,
		passwordExp: passwordExp,
	}
}

func (u *UserHandler) RegisterRoutes(server *gin.Engine) {
	ug := server.Group("/users")
	ug.GET("/profile", u.Profile)
	ug.POST("/signup", u.SingUp)
	ug.POST("/login", u.Login)
	ug.POST("/edit", u.Edit)
}

func (u *UserHandler) SingUp(ctx *gin.Context) {
	type SignUpReq struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}

	var res SignUpReq
	//Bind 方法会根据 content-type 来解析你的数据到res里
	//解析错了会返回400
	if err := ctx.Bind(&res); err != nil {
		return
	}
	if res.Password != res.ConfirmPassword {
		ctx.String(http.StatusOK, "两次密码输入不一致")
		return
	}

	ok, err := u.emailExp.MatchString(res.Email)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "系统错误")
		return
	}
	if !ok {
		ctx.String(http.StatusOK, "你的邮箱格式不正确")
		return
	}
	ok, err = u.passwordExp.MatchString(res.Password)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "系统错误")
		return
	}
	if !ok {
		ctx.String(http.StatusOK, "密码必须大于8位，包含数字，特殊字符")
		return
	}

	ctx.String(http.StatusOK, "注册成功")
	fmt.Printf("%v", res)
}
func (u *UserHandler) Login(ctx *gin.Context) {

}
func (u *UserHandler) Edit(ctx *gin.Context) {

}
func (u *UserHandler) Profile(ctx *gin.Context) {

}
