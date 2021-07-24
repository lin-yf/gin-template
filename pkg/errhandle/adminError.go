package errhandle

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func setCookie(c *gin.Context, key, value string) *gin.Context {
	c.SetCookie(key, value, 100, "/", "", false, false)
	return c
}

func ResponseError(c *gin.Context, status string) *gin.Context {
	res, _ := json.Marshal([]string{"aaa:123", "bbb:456"})
	setCookie(c, "notice", string(res))
	setCookie(c, "notice_type", status)
	// setCookie(c, "notice_highlight", status)
	return c
}
