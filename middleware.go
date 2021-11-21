package main

import (
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"pgd/global"
	"time"
)

func middleware() (*jwt.GinJWTMiddleware, error) {
	// the jwt middleware
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "pgd",
		Key:         []byte("wdnmd"),
		Timeout:     time.Hour * 240,
		MaxRefresh:  time.Hour * 240,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims { //前端能解码的数据
			if v, ok := data.(*global.User); ok {
				return jwt.MapClaims{
					identityKey: v.ID,
					"name":      v.Name,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} { //登陆中间件，私有变量
			claims := jwt.ExtractClaims(c)
			fmt.Println(claims)
			c.Set("uid", claims["id"])
			c.Set("name", claims["name"])

			return &global.User{}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) { //验证
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userID := loginVals.ID
			password := loginVals.Password

			user := global.User{}
			global.DB.Where(`id = ? AND password = ?`, userID, password).Find(&user)
			if user.ID != "" {
				return &user, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool { //鉴权
			if _, ok := data.(*global.User); ok {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) { //登录失败或者其他
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},

		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})
}
