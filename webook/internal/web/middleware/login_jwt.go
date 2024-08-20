package middleware

import (
	"encoding/gob"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
	"time"
)

type LoginJWTMiddlewareBuilder struct {
	paths []string
}

func NewLoginJWTMiddlewareBuilder() *LoginJWTMiddlewareBuilder {
	return &LoginJWTMiddlewareBuilder{}
}

func (l *LoginJWTMiddlewareBuilder) IgnorePaths(path string) *LoginJWTMiddlewareBuilder {
	l.paths = append(l.paths, path)
	return l
}

func (l *LoginJWTMiddlewareBuilder) Build() gin.HandlerFunc {
	//Go的方式编码解码
	gob.Register(time.Now())
	return func(ctx *gin.Context) {
		//不需要登录校验
		for _, path := range l.paths {
			if ctx.Request.URL.Path == path {
				return
			}
		}
		// JWT 来校验
		tokenHeader := ctx.GetHeader("Authorization")
		if tokenHeader == "" {
			//没有登录
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		segs := strings.Split(tokenHeader, " ")
		if len(segs) != 2 {
			//没登录，有人在瞎搞
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenStr := segs[1]
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte("rvPJQMnIiS9oQ447JedsZapn2PydxUgo"), nil
		})

		if err != nil {
			//没有登录
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// err 为 nil， token 不为 nil
		if token == nil || !token.Valid {
			//没有登录
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
