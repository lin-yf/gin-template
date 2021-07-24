package models

import (
	"fmt"

	"go-template/pkg/log"

	"go-template/setting"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Global Account
var (
	DB *gorm.DB
)

// DB : Variable for interacting with database
func init() {
	log.Info("db init")
}

// GormInit init gorm ORM.
func GormInit() {
	log.Info("gorm init")
	var db *gorm.DB
	var openErr error
	switch setting.Conf.Database.Type {
	case "UNSET", "sqlite", "sqlite3":
		dbURL := setting.Conf.Database.FilePath
		db, openErr = gorm.Open(sqlite.Open(dbURL), &gorm.Config{})
	case "postgres":
		dbURL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
			setting.Conf.Database.Host,
			setting.Conf.Database.User,
			setting.Conf.Database.Password,
			setting.Conf.Database.DBName,
			setting.Conf.Database.Port,
		)
		db, openErr = gorm.Open(postgres.Open(dbURL), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Warn),
		})
	default:
		log.Fatalf("未设置数据库类型：%s", setting.Conf.Database.Type)
	}
	if openErr != nil {
		log.Fatalf("连接数据库错误：%s", openErr)
	}
	sqlDB, sqlerr := db.DB()
	if sqlerr != nil {
		log.Panicf("sqlDB错误：%s\n", sqlerr.Error())
	}
	// 表前缀
	// gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	// 	return conf.DatabaseConfig.TablePrefix + defaultTableName
	// }
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	// 初始化表
	db.AutoMigrate(&Account{})
	if db.Error != nil {
		log.Errorf("初始化db错误：%s\n", db.Error.Error())
		return
	}
	DB = db
}

// GetDB 获取db
func GetDB() *gorm.DB {
	return DB
}
