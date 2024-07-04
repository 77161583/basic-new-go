package web

import (
	"basic-new-go/webook/internal/domain"
	"basic-new-go/webook/internal/service"
	"errors"
	"fmt"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

// UserHandler 定义所有跟user用户有关的路由
type UserHandler struct {
	svc         *service.UserService
	emailExp    *regexp.Regexp
	passwordExp *regexp.Regexp
	birthdayExp *regexp.Regexp
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	const (
		emailRegex            = "^[a-z0-9_\\.\\-\\+]+@[a-z0-9\\-]+(\\.[a-z0-9\\-]+)*$"
		passwordRegex         = `^(?=.*[A-Z])(?=.*[a-z])(?=.*[0-9])(?=.*[!@#$%^&*()_+=-])[A-Za-z0-9!@#$%^&*()_+=-]{8,}$`
		birthdayRegExpPattern = `^\d{4}-\d{2}-\d{2}$`
	)
	emailExp := regexp.MustCompile(emailRegex, regexp.None)
	passwordExp := regexp.MustCompile(passwordRegex, regexp.None)
	birthdayRegExp := regexp.MustCompile(birthdayRegExpPattern, regexp.None)
	return &UserHandler{
		svc:         svc,
		emailExp:    emailExp,
		passwordExp: passwordExp,
		birthdayExp: birthdayRegExp,
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
		fmt.Println(err)
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
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "注册成功",
	})
	//ctx.String(http.StatusOK, "注册成功")
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
	sess.Options(sessions.Options{
		//生产环境需要设置 https
		//Secure:   true,
		//HttpOnly: true,
		MaxAge: 60,
	})
	sess.Save()

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "登录成功",
	})
	//ctx.String(http.StatusOK, "登录成功")

	return
}
func (u *UserHandler) LoginOut(ctx *gin.Context) {
	sess := sessions.Default(ctx)
	//设置session的值
	sess.Options(sessions.Options{
		MaxAge: 1, //cookie过期时间
	})
	sess.Save()
	ctx.String(http.StatusOK, "退出登录成功")
	return
}
func (u *UserHandler) Edit(ctx *gin.Context) {
	type EditReq struct {
		Id              int64  `json:"id"`
		NickName        string `json:"nickName"`
		Birthday        string `json:"birthday"`
		PersonalProfile string `json:"personalProfile"`
	}
	var req EditReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.String(http.StatusBadRequest, "参数绑定失败")
		return
	}

	// 打印接收到的请求参数
	fmt.Printf("Received Edit Request: %+v\n", req)

	//参数校验
	if req.NickName != "" {
		if len(req.NickName) < 6 || len(req.NickName) > 30 {
			ctx.String(http.StatusOK, "昵称大小需要保持2~10个汉字")
			return
		}
	}
	if req.Birthday != "" {
		isBirthday, _ := u.birthdayExp.MatchString(req.Birthday)
		if !isBirthday {
			ctx.String(http.StatusOK, "生日日期格式不正确，应为YYYY-MM-DD格式")
			return
		}
	}
	if req.PersonalProfile == "" {
		ctx.String(http.StatusOK, "个人简介不能为空")
		return
	}

	err := u.svc.Edit(ctx, domain.User{
		Id:              req.Id,
		NickName:        strings.TrimSpace(req.NickName),
		Birthday:        strings.TrimSpace(req.Birthday),
		PersonalProfile: strings.TrimSpace(req.PersonalProfile),
	})
	if errors.Is(err, service.ErrUserIdNotFund) {
		ctx.String(http.StatusOK, "用户不存在")
		return
	}
	if err != nil {
		// 打印错误信息
		fmt.Printf("Service Edit Error: %v\n", err)
		ctx.String(http.StatusOK, "更新失败")
		return
	}
	ctx.String(http.StatusOK, "更新成功")
}
func (u *UserHandler) Profile(ctx *gin.Context) {
	// 从 URL 查询参数中获取 id
	idStr := ctx.Query("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.String(http.StatusBadRequest, "无效的用户 ID")
		return
	}
	userData, err := u.svc.Profile(ctx, id)
	if err != nil {
		if errors.Is(err, service.ErrUserIdNotFund) {
			ctx.String(http.StatusNotFound, "用户不存在")
			return
		}
		// 其他错误处理
		ctx.String(http.StatusInternalServerError, "获取用户数据失败")
		return
	}
	ctx.JSON(http.StatusOK, userData)
}
