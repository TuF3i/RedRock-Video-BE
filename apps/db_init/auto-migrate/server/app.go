package server

import (
	"LiveDanmu/apps/db_init/auto-migrate/core/dao"
	"LiveDanmu/apps/shared/config"
	"os"

	"gitee.com/liumou_site/logger"
)

var l *logger.LocalLogger

func RunInitDB() {
	l = logger.NewLogger(1)
	l.Modular = "db_init-on-create"
	l.Info("Starting AutoMigrate Process...")

	// 初始化配置文件
	conf, err := config.LoadDBInitConfig()
	if err != nil {
		l.Error("Load Configuration Error: %v", err.Error())
		os.Exit(1)
	}

	// 初始化dao
	dao, err := dao.GetDao(conf)
	if err != nil {
		l.Error("Init Dao Error: %v", err.Error())
		os.Exit(1)
	}

	// 开始数据库迁移
	err = dao.AutoMigrate()
	if err != nil {
		l.Error("AutoMigrate Fail: %v", err.Error())
	}

	l.Info("AutoMigrate Success!")
}
