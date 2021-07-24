package captchas

import (
	"bytes"
	"net/http"
	"time"

	"github.com/dchest/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Captcha 设置图片验证码
func Captcha(c *gin.Context, length ...int) {
	l := captcha.DefaultLen
	w, h := 107, 36
	if len(length) == 1 {
		l = length[0]
	}
	if len(length) == 2 {
		w = length[1]
	}
	if len(length) == 3 {
		h = length[2]
	}
	captchaID := captcha.NewLen(l)
	session := sessions.Default(c)
	// log.Infof("captchaID:%s", captchaID)
	session.Set("captcha", captchaID)
	_ = session.Save()
	_ = Serve(c.Writer, c.Request, captchaID, ".png", "zh", false, w, h)
}

// CaptchaVerify 图片验证码验证
func CaptchaVerify(c *gin.Context, code string) bool {
	// session := sessions.Default(c)
	// if captchaID := utils.GetSession(c, "captcha"); captchaID != nil {
	// 	// session.Delete("captcha")
	// 	// utils.DelSession(c, "captcha")
	// 	if captcha.VerifyString(captchaID.(string), code) {
	// 		return true
	// 	} else {
	// 		return false
	// 	}
	// } else {
	// 	return false
	// }
	return true
}

func Serve(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, height int) error {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	var content bytes.Buffer
	switch ext {
	case ".png":
		w.Header().Set("Content-Type", "image/png")
		_ = captcha.WriteImage(&content, id, width, height)
	case ".wav":
		w.Header().Set("Content-Type", "audio/x-wav")
		_ = captcha.WriteAudio(&content, id, lang)
	default:
		return captcha.ErrNotFound
	}

	if download {
		w.Header().Set("Content-Type", "application/octet-stream")
	}
	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
	return nil
}
