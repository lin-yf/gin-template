package main

import (
	"fmt"
	"go-template/pkg/crontab"
	"go-template/routers"
	"go-template/setting"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	setting.Init()
	// models.GormInit()
	crontab.Init()
}

func main() {
	if os.Getenv("MODE") == "prod" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	routersInit := routers.InitRouter()
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	addr := fmt.Sprintf("%s:%s", host, port)
	server001 := &http.Server{
		Addr:         addr,
		Handler:      routersInit,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	log.Printf("[info] start http server listening %s", addr)
	server001.ListenAndServe()
}
