package app

import (
	"context"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"go-urlsaver/internal/config"
	"go-urlsaver/internal/handler"
	"go-urlsaver/internal/repository"
	"go-urlsaver/internal/repository/postgres"
	"go-urlsaver/internal/server"
	"go-urlsaver/internal/service"
	"os"
	"os/signal"
	"syscall"
)

func Run(configPath string) {
	cfg, err := config.Init(configPath)

	if err != nil {
		logrus.Fatalf("error occurred when reading congif: %s", err.Error())
	}

	db, err := postgres.NewPostgresDB(cfg.DB)
	if err != nil {
		logrus.Fatalf("error occurred when connecting to db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)

	go func() {
		if err := srv.Run(cfg, handlers.Init()); err != nil {
			logrus.Fatalf("error occurred while running http server: %s", err.Error())
		}
	}()

	logrus.Println("Url Saver App Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Println("Url Saver App Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occurred while shutting down server: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occurred while closing database: %s", err.Error())
	}
}
