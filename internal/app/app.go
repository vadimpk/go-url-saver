package app

import (
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"go-urlsaver/internal/config"
	"go-urlsaver/internal/handler"
	"go-urlsaver/internal/repository"
	"go-urlsaver/internal/repository/postgres"
	"go-urlsaver/internal/server"
	"go-urlsaver/internal/service"
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
	err = srv.Run(cfg, handlers.Init())
	if err != nil {
		logrus.Fatalf("error occurred while running http server: %s", err.Error())
	}

}
