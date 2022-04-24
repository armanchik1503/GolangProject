package main

import (
	todo "github.com/armanchik1503/projectGolang"
	handler2 "github.com/armanchik1503/projectGolang/pkg/handler"
	"github.com/armanchik1503/projectGolang/pkg/repository"
	"github.com/armanchik1503/projectGolang/pkg/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	//if err := initConfig(); err != nil {
	//	logrus.Fatalf("error initializing configs: %s", err.Error())
	//}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("SSL_MODE"),
	})
	if err != nil {
		logrus.Fatalf("Failed to initialize database: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler2.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(os.Getenv("SERVER_PORT"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Error ocured while running http server: %s", err.Error())
	}
}

//func initConfig() error {
//	viper.AddConfigPath("configs")
//	viper.SetConfigName("config")
//
//	return viper.ReadInConfig()
//}
