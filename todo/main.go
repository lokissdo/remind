package main

import (
	"log"
	"net"
	"remind/pkg/logger"
	"remind/todo/api/middleware"
	"remind/todo/common/configs"
	"remind/todo/gin_service"
	"remind/todo/service"

	// _ "remind/todo/repository"
	//"todo
	"os"
	api "remind/todo/api"
	pb "remind/todo/pb"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
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

// func main() {
// 	setupLog()

// 	r := setupRouter()
// 	r.Run(fmt.Sprintf(":%d", configs.Yaml.Port))
// }

func main() {
    port := ":5005"
    lis, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatalf("Failed to listen on port %v: %v", port, err)
    }

	grpcLogger := grpc.UnaryInterceptor(logger.GrpcLogger)

	s := grpc.NewServer(grpcLogger)
    pb.RegisterTodoServiceServer(s, &service.TodoServiceServer{})
	// Run reminder cron job
	go gin_service.RunReminderCronJob()
    log.Printf("Server listening at %v", lis.Addr())
    if err := s.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}