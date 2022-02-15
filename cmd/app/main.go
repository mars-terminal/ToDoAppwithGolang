package main

import (
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"myToDoApp/internal/handler"
	"myToDoApp/internal/repository"
	"myToDoApp/internal/service/server"
	"myToDoApp/internal/service/service"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

var log = logrus.WithField("package", "main")

func main() {
	if err := InitConfig(); err != nil {
		log.Fatalf("error initializing configs : %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %v", err)
	}

	var opts = struct {
		DBHost     string
		DBPort     string
		DBUser     string
		DBPassword string
		DBName     string
		DBSSLMode  string

		HttpPORT string
	}{
		DBHost:     viper.GetString("db.host"),
		DBPort:     viper.GetString("db.port"),
		DBUser:     viper.GetString("db.user"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     viper.GetString("db.dbname"),
		DBSSLMode:  viper.GetString("db.sslmode"),
		HttpPORT:   viper.GetString("http.port"),
	}

	db, err := repository.PostgresNewDB(repository.Config{
		Host:     opts.DBHost,
		Port:     opts.DBPort,
		UserName: opts.DBUser,
		Password: opts.DBPassword,
		DBName:   opts.DBName,
		SSLMode:  opts.DBSSLMode,
	})

	if err != nil {
		log.WithError(err).Fatal("failed to initialize db")
	}
	log.Info("db connected")

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	if err := srv.Run(opts.HttpPORT, handlers.InitRoutes()); err != nil {
		log.WithError(err).Fatal("error occured while running http server")
	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
