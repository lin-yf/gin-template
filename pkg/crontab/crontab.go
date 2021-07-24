package crontab

import (
	"go-template/pkg/log"

	"github.com/robfig/cron/v3"
)

// Cron 定时任务
var Cron *cron.Cron

// Reload 重新启动定时任务
func Reload() {
	if Cron != nil {
		Cron.Stop()
	}
	Init()
}

// Init 初始化
func Init() {
	log.Infof("初始化定时任务...")
	Cron := cron.New()
	_, err := Cron.AddFunc("* * * * *", func() {
		// fmt.Println("Every minute on the half hour")
	})
	if err != nil {
		log.Warnf("定时任务添加失败")
	}
	Cron.Start()
}
