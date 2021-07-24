package cookies

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func getDomainName() string {
	domain := ""
	if os.Getenv("MODE") == "dev" {
		domain = ""
	}
	return domain
}

// Login 登录控制
func setLogin(c *gin.Context) (int, error) {
	// maxAge, err := strconv.ParseUint(os.Getenv("COOKIE_MAX_AGE"), 10, 8)
	// if err != nil {
	// 	return http.StatusBadRequest, err
	// }
	maxAge := 10000
	if c.PostForm("remember_me") == "remember" {
		maxAge = 365 * 24 * 3600
	}
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "session",
		Value:   "",
		Domain:  getDomainName(),
		Path:    "/",
		Expires: time.Now().AddDate(-1, -1, -1),
	})
	c.SetCookie("session", "99", maxAge, "/", "", false, true)
	return http.StatusOK, nil
}
