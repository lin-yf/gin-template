package utils

import (
	"go-template/pkg/log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// SetSession 设置session
func SetSession(c *gin.Context, list map[string]interface{}) {
	s := sessions.Default(c)
	for key, value := range list {
		s.Set(key, value)
	}
	err := s.Save()
	if err != nil {
		log.Warnf("设置session 失败：%s", err)
	}
}

// GetSession 获取session
func GetSession(c *gin.Context, key string) interface{} {
	s := sessions.Default(c)
	return s.Get(key)
}

// DelSession 删除session
func DelSession(c *gin.Context, key string) {
	s := sessions.Default(c)
	s.Delete(key)
	s.Save()
}

// IsLogin 是否登录
func IsLogin(c *gin.Context) bool {
	username := GetSession(c, "username")
	if username == nil {
		return false
	}
	return true
}
