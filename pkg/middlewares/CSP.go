package middlewares

import "github.com/gin-gonic/gin"

// CSP set Content Security Policy http header
func CSP() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Security-Policy", "default-src 'self'; img-src * data:; media-src *; style-src 'self' maxcdn.bootstrapcdn.com fonts.googleapis.com 'unsafe-inline'; script-src 'self' 'unsafe-inline' 'unsafe-eval' a-ads.com *.a-ads.com; font-src 'self' fonts.gstatic.com maxcdn.bootstrapcdn.com; child-src ad.a-ads.com a-ads.com *.a-ads.com")
		c.Next()
	}
}
