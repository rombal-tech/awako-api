package main

import (
	"alvile-api/pkg/handler"
	"alvile-api/pkg/repository"
	"alvile-api/pkg/service"
	"alvile-api/server"
	"github.com/execaus/exloggo"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"

	"github.com/spf13/viper"
	"os"
)

func main() {
	if err := exloggo.SetParameters(&exloggo.Parameters{
		Directory: "logs",
	}); err != nil {
		log.Fatal(err.Error())
	}

	if err := initConfig(); err != nil {
		exloggo.Fatalf("error initializing configs: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		exloggo.Warningf("error loading env variables: %s, use os env", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("postgres.host"),
		Port:     viper.GetString("postgres.port"),
		Username: viper.GetString("postgres.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("postgres.dbname"),
		SSLMode:  viper.GetString("postgres.sslmode"),
	})
	if err != nil {
		exloggo.Fatalf("failed to initialize db: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewService(services)
	var serverInstance server.Server
	if err := serverInstance.Run(viper.GetString("server.port"), handlers.InitRoutes()); err != nil {
		exloggo.Fatalf("error ocured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config.local")
	return viper.ReadInConfig()
}
