package controllers

import (
	"net/http"

	"go-template/pkg/errhandle"

	"github.com/gin-gonic/gin"
)

// SuccessJSON 成功
func SuccessJSON(c *gin.Context, msg string, data interface{}) *gin.Context {
	c.JSON(http.StatusOK, errhandle.Response{
		Code:   100,
		Msg:    msg,
		Result: data,
	})
	return c
}
