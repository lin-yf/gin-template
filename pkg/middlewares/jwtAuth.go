package middlewares

import (
	"errors"
	"net/http"

	"go-template/pkg/errhandle"
	"go-template/pkg/jwt"

	"github.com/gin-gonic/gin"
)

var (
	SignKey string = "owl" // 签名信息应该设置成动态从库中获取
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取token
		token := c.Request.Header.Get("X-Owl-Token")
		if token == "" {
			c.JSON(http.StatusUnauthorized, errhandle.ErrWithData(errhandle.AppCheckLogin, "token 不存在", errors.New("token 不存在"), gin.H{}))
			c.Abort()
			return
		}
		// 初始化一个JWT对象实例，并根据结构体方法来解析token
		j := jwt.NewJWT()
		// 解析token中包含的相关信息(有效载荷)
		_, claims, err := j.ParseToken(token)
		if err != nil {
			// token过期
			if err == errhandle.TokenExpired {
				c.JSON(http.StatusUnauthorized, errhandle.ErrWithData(errhandle.AppCheckLogin, "token过期", errors.New("无此权限"), gin.H{
					"isLogin": true,
				}))
				c.Abort()
				return
			}
			// 其他错误
			c.JSON(http.StatusUnauthorized, errhandle.ErrWithData(errhandle.AppUnknownErr, "token过期", errors.New("无此权限"), gin.H{}))
			c.Abort()
			return
		}
		// 将解析后的有效载荷claims重新写入gin.Context引用对象中
		c.Set("claims", claims)
		c.Next()
	}
}
