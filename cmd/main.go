package main

import (
	"alvile-api/pkg/handler"
	"alvile-api/pkg/repository"
	"alvile-api/pkg/service"
	"alvile-api/server"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewService(services)
	var serverInstance server.Server
	if err := serverInstance.Run("8080", handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error ocured while running http server: %s", err.Error())
	}

}
