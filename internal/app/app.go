package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/BabyJhon/backend-trainee-assignment-2024/configs"
	"github.com/BabyJhon/backend-trainee-assignment-2024/internal/handlers"
	"github.com/BabyJhon/backend-trainee-assignment-2024/internal/repo"
	"github.com/BabyJhon/backend-trainee-assignment-2024/internal/service"
	"github.com/BabyJhon/backend-trainee-assignment-2024/pkg/httpserver"
	"github.com/BabyJhon/backend-trainee-assignment-2024/pkg/postgres"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Run() {
	//logrus
	logrus.SetFormatter(new(logrus.JSONFormatter))

	//Configs
	if err := configs.InitConfig(); err != nil {
		logrus.Fatalf("error initialization configs: %s", err.Error())
	}

	//.env
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env vars: %s", err.Error())
	}

	//DB
	pool, err := postgres.NewPG(context.Background(), postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("failed init db: %s", err.Error())
	}

	defer pool.Close()

	// Repositories
	repos := repo.NewRepository(pool)

	// Service
	services := service.NewService(repos)

	// Handlers
	handlers := handlers.NewHandler(services)

	//HTTP server

	srv := new(httpserver.Server)

	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != http.ErrServerClosed {
			logrus.Fatalf("error occured while running server: %s", err.Error())
		}
	}()

	logrus.Print("avito-tech-backend-trainee-assignment-2024 API started")

	//gracefull shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("shutting down")
	if err := srv.ShutDown(context.Background()); err != nil {
		logrus.Errorf("error while server shutting down: %s", err.Error())
	}

}
