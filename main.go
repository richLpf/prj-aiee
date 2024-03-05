package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"prj-aiee/config"
	"prj-aiee/router"
	_ "prj-aiee/service/apis"
	"prj-aiee/service/logger"
	"prj-aiee/service/mysql"
)

func main() {
	configFile := flag.String("c", "./config/config.yml", "yml config")
	flag.Parse()
	if *configFile == "" {
		fmt.Println("invalid config file")
		os.Exit(0)
	}
	// 加载配置文件
	config.LoadConfig(*configFile)
	logger.Logger.Info("Service Start")
	db := mysql.NewMySQL(context.Background(), &config.Cfg.MySQL)
	db.WithLogger(logger.Logger)
	db.ConnDB()
	router.Init()
}
