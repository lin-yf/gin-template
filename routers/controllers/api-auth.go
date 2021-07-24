package controllers

import (
	"net/http"

	"go-template/models"
	"go-template/pkg/errhandle"
	"go-template/pkg/jwt"
	"go-template/pkg/log"
	"go-template/pkg/utils"
	"go-template/pkg/validator/user"

	"github.com/gin-gonic/gin"
)

func LoginPost(c *gin.Context) {
	b := user.LoginForm{}
	if err := c.ShouldBind(&b); err != nil {
		log.Errorf("登录校验失败", err)
		c.JSON(http.StatusUnauthorized, errhandle.Err(errhandle.AppCheckLogin, "字端校验失败", err))
		return
	}
	var loginUser models.Account
	err := loginUser.FindUserByName(b.Username)
	if err != nil {
		log.Errorf("用户名或密码错误", err)
		c.JSON(http.StatusUnauthorized, errhandle.ErrWithData(errhandle.AppCheckLogin, "用户不存在", err, gin.H{}))
		return
	}
	if err := utils.CheckPassword(b.Password, loginUser.Password); err != nil {
		log.Errorf("用户名或密码错误", err)
		c.JSON(http.StatusUnauthorized, errhandle.ErrWithData(errhandle.AppCheckLogin, "登录校验失败", err, gin.H{}))
		return
	}
	token, err := jwt.GenerateToken(loginUser.Username)
	if err != nil {
		log.Errorf("token生成失败", err)
		c.JSON(http.StatusUnauthorized, errhandle.ErrWithData(errhandle.AppCheckLogin, "token生成失败", err, gin.H{}))
		return
	}
	SuccessJSON(c, "保存成功", gin.H{
		"token":     token,
		"avatar":    "https://gw.alipayobjects.com/zos/rmsportal/jZUIxmJycoymBprLOUbT.png",
		"status":    1,
		"creatorId": "admin",
		"roleId":    "admin",
		"lang":      "zh-CN",
	})
}

func GetCurrentNav(c *gin.Context) {
	SuccessJSON(c, "success", gin.H{
		"avatar":    "https://gw.alipayobjects.com/zos/rmsportal/jZUIxmJycoymBprLOUbT.png",
		"status":    1,
		"creatorId": "admin",
		"roleId":    "admin",
		"lang":      "zh-CN",
	})
}
