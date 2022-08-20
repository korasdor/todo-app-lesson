package main

import (
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
	if err := utils.InitConfigs(); err != nil {
		logrus.Println("Error occurred while reading env file, might fallback to OS env config")
	}
	gin.SetMode(utils.GetEnvVar("GIN_MODE"))

	db, err := repository.NewMysqlDB(utils.GetEnvVar("DB_DATASOURCE_NAME"))

	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	logrus.Println("Starting run server...")

	srv := new(core.Server)
	if err := srv.Run(utils.GetEnvVar("PORT"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}

}
