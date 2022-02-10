/**
@author ZhengHao Lou
Date    2022/02/03
*/
package main

import (
	"flag"
	"fmt"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/internal/model"
	"github.com/go-programming-tour-book/blog-service/internal/routers"
	"github.com/go-programming-tour-book/blog-service/pkg/logger"
	"github.com/go-programming-tour-book/blog-service/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	port    string
	runMode string
	config  string

	isVersion    bool
	buildTime    string
	buildVersion string
	gitCommitID  string
)

func init() {
	var err error
	err = setupFlag()
	if err != nil {
		log.Fatalf("init.setupFlag error: %v\n", err)
	}
	err = setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting error: %v\n", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine error: %v\n", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger error: %v\n", err)
	}
}

func setupFlag() error {
	flag.StringVar(&port, "port", "", "startup port")
	flag.StringVar(&runMode, "mode", "", "startup mode")
	flag.StringVar(&config, "config", "/configs", "config file path")
	flag.BoolVar(&isVersion, "version", false, "compile information")
	flag.Parse()
	return nil
}

func setupSetting() error {
	setting, err := setting.NewSetting(strings.Split(config, ",")...)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Jwt", &global.JwtSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	global.JwtSetting.Expire *= time.Second
	if port != "" {
		global.ServerSetting.HttpPort = port
	}
	if runMode != "" {
		global.ServerSetting.RunMode = runMode
	}
	return nil
}

func setupDBEngine() (err error) {
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	return
}

func setupLogger() (err error) {
	fileName := global.AppSetting.LogSavePath + "/" +
		global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return
}

// @title blog
// @version 1.0
// @description Go programming tour
// @termsOfService https://github.com/go-programming-tour-book
func main() {
	if isVersion {
		fmt.Printf("build_time: %s\n", buildTime)
		fmt.Printf("build_version: %s\n", buildVersion)
		fmt.Printf("git_commit_id: %s\n", gitCommitID)
		return
	}

	global.Logger.Infof("%s: go-programming-tour-book/%s",
		"my", "blog-service")
	r := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        r,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
