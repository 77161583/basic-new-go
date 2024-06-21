package web

import (
	"basic-new-go/webook/internal/domain"
	"basic-new-go/webook/internal/service"
	"errors"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserHandler 定义所有跟user用户有关的路由
type UserHandler struct {
	svc         *service.UserService
	emailExp    *regexp.Regexp
	passwordExp *regexp.Regexp
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	const (
		emailRegex    = "^[a-z0-9_\\.\\-\\+]+@[a-z0-9\\-]+(\\.[a-z0-9\\-]+)*$"
		passwordRegex = `^(?=.*[A-Z])(?=.*[a-z])(?=.*[0-9])(?=.*[!@#$%^&*()_+=-])[A-Za-z0-9!@#$%^&*()_+=-]{8,}$`
	)
	emailExp := regexp.MustCompile(emailRegex, regexp.None)
	passwordExp := regexp.MustCompile(passwordRegex, regexp.None)
	return &UserHandler{
		svc:         svc,
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

	err = u.svc.SignUp(ctx, domain.User{
		Email:    res.Email,
		Password: res.Password,
	})

	if errors.Is(err, service.ErrUserDuplicateEmail) {
		ctx.String(http.StatusOK, "重复邮箱，请换一个")
		return
	}

	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	ctx.String(http.StatusOK, "注册成功")
	//fmt.Printf("%v", res)
}
func (u *UserHandler) Login(ctx *gin.Context) {
	type LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req LoginReq
	if err := ctx.Bind(&req); err != nil {
		return
	}
	user, err := u.svc.Login(ctx, req.Email, req.Password)
	if errors.Is(err, service.ErrInvalidUserOrPassword) {
		ctx.String(http.StatusOK, "用户名或密码不对")
		return
	}
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}

	//登录/步骤2
	//登陆成功之后设置session
	sess := sessions.Default(ctx)
	//设置session的值
	sess.Set("userId", user.Id)
	sess.Save()
	ctx.String(http.StatusOK, "登录成功")

	return
}
func (u *UserHandler) Edit(ctx *gin.Context) {

}
func (u *UserHandler) Profile(ctx *gin.Context) {
	ctx.String(http.StatusOK, "i am profile")
}
