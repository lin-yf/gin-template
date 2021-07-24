package middlewares

import (
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
)

// Store session存储
var Store memstore.Store

// Session 初始化session
func Session(secret string) gin.HandlerFunc {
	// Redis设置不为空，且非测试模式时使用Redis
	// if conf.RedisConfig.Server != "" && gin.Mode() != gin.TestMode {
	// 	var err error
	// 	Store, err = redis.NewStoreWithDB(10, conf.RedisConfig.Network, conf.RedisConfig.Server, conf.RedisConfig.Password, conf.RedisConfig.DB, []byte(secret))
	// 	if err != nil {
	// 		log.Errorf("无法连接到 Redis：%s", err)
	// 	}

	// 	log.Infof("已连接到 Redis 服务器：%s", conf.RedisConfig.Server)
	// } else {
	// 	Store = memstore.NewStore([]byte(secret))
	// }

	// Store = memstore.NewStore([]byte(secret))
	Store := cookie.NewStore([]byte(secret))
	// sessionNames := []string{"se", "on"}
	// r.Use(sessions.SessionsMany(sessionNames, store))
	Store.Options(sessions.Options{
		HttpOnly: true,
		MaxAge:   7 * 86400,
		Path:     "/",
		Secure:   os.Getenv("MODE") == "prod",
	})
	return sessions.Sessions("owl", Store)
}
