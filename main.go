package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/korasdor/todo-app/core"
	"github.com/korasdor/todo-app/handler"
	"github.com/korasdor/todo-app/repository"
	"github.com/korasdor/todo-app/service"
	"github.com/korasdor/todo-app/utils"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	logrus.SetReportCaller(true)
	if err := utils.InitConfigs(); err != nil {
		logrus.Println("Error occurred while reading env file, might fallback to OS env config")
	}
	gin.SetMode(utils.GetEnvVar("GIN_MODE"))

	db, err := repository.NewMysqlDB(utils.GetEnvVar("DB_DATASOURCE_NAME"))

	if err != nil {
		logrus.Fatalf("failed to initialize db %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(core.Server)
	go func() {
		if err := srv.Run(utils.GetEnvVar("PORT"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("TodoApp Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	logrus.Print("TodoApp Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on database connection close: %s", err.Error())
	}

}
