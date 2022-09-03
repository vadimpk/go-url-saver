package app

import (
	_ "github.com/lib/pq"
	"go-urlsaver/internal/config"
	"go-urlsaver/internal/handler"
	"go-urlsaver/internal/repository"
	"go-urlsaver/internal/server"
	"go-urlsaver/internal/service"
	"log"
)

func Run(configPath string) {
	cfg, err := config.Init(configPath)

	if err != nil {
		log.Fatalf("error occurred when reading congif: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "lz921skm0001p",
		DBName:   "postgres",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalf("error occurred when connecting to db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	err = srv.Run(cfg, handlers.Init())
	if err != nil {
		log.Fatalf("error occurred while running http server: %s", err.Error())
	}

}
