package main

import (
	configs "alvile-api/config"
	"alvile-api/models"
	"alvile-api/repository"
	"alvile-api/server"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	var serverInstance server.Server
	env := models.LoadEnv()
	config := configs.LoadConfig()
	database := repository.NewBusinessDatabase(env, config)

	if err := serverInstance.Run("8080"); err != nil {

	}
}
