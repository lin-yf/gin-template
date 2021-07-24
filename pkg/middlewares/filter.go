package middlewares

import (
	"errors"
	"net/http"

	"go-template/pkg/utils"

	"go-template/pkg/errhandle"

	"github.com/gin-gonic/gin"
)

// AuthFilter 登陆过滤
func AuthFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !utils.IsLogin(c) {
			c.Abort()
			c.Redirect(http.StatusFound, "/admin/login")
			return
		}
		c.Next()
	}
}

// ApiAuthFilter 接口登录过滤
func ApiAuthFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !utils.IsLogin(c) {
			c.Abort()
			c.JSON(http.StatusUnauthorized, errhandle.ErrWithData(errhandle.AppCheckLogin, "登录校验失败", errors.New("无此权限"), gin.H{
				"isLogin": true,
			}))
			return
		}
		c.Next()
	}
}
