package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Online 全局拦截
func Online() gin.HandlerFunc {
	return func(c *gin.Context) {
		// if cookie, err := c.Request.Cookie("owl:on"); err != nil {
		// 	sessions.Set("on", )
		// } else {

		// }
		session := sessions.DefaultMany(c, "on")
		var count int
		v := session.Get("on")
		if v == nil {
			count = 0
		} else {
			count = 10
		}
		session.Set("on", count)
		c.Next()
	}
}
