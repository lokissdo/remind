package main

import (
	"todo/common/configs"
	"todo/api/middleware"

	// _ "todo/repository"
	//"todo
	api "todo/api"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.LogRequest())
	api.Register(r)
	return r
}
func setupLog() {
	if configs.Yaml.Log.File == "" {
		logrus.SetOutput(os.Stdout)
	} else {
		src, err := os.OpenFile(configs.Yaml.Log.File, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		//Write to file

		if err != nil {
			logrus.WithFields(logrus.Fields{
				"error": err,
			}).Error("Init log error")
			panic(err)
		}

		//Set output
		logrus.SetOutput(src)
	}

	//Set log level
	if configs.Yaml.Log.Level == "debug" {
		logrus.SetLevel(logrus.DebugLevel)
	} else if configs.Yaml.Log.Level == "warn" {
		logrus.SetLevel(logrus.WarnLevel)
	} else if configs.Yaml.Log.Level == "error" {
		logrus.SetLevel(logrus.ErrorLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	//Format log
	logrus.SetFormatter(&logrus.JSONFormatter{})

}

func main() {
	setupLog()
	r := setupRouter()
	r.Run(fmt.Sprintf(":%d", configs.Yaml.Port))
}
