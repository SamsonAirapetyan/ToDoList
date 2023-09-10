package main

import (
	"fmt"
	"github.com/SamsonAirapetyan/todo-app"
	"github.com/SamsonAirapetyan/todo-app/pkg/handler"
	"github.com/SamsonAirapetyan/todo-app/pkg/repository"
	"github.com/SamsonAirapetyan/todo-app/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error init config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error with loading password %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	fmt.Println("After postgreess")
	if err != nil {
		logrus.Fatalf("failed conection with BD %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	// handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatal(err)
	}

}

// служит для получения данных с конфиг, для этого нужна библиотека viper
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
