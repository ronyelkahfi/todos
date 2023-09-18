package main

import (
	httpAPI "github.com/ronyelkahfi/todos/internal/delivery/rest/api"
	todoRepo "github.com/ronyelkahfi/todos/internal/domain/todo/repository"
	todoDB "github.com/ronyelkahfi/todos/internal/domain/todo/repository/db"
	"github.com/ronyelkahfi/todos/internal/service"
	"github.com/ronyelkahfi/todos/internal/utils"
	"github.com/ronyelkahfi/todos/pkg/http"
)
func server() error {
	// Get config.
	cfg, err := getConfig()
	if err != nil {
		return err
	}
	utils.Info("config initialized")
	// Init db.
	db, err := newDB(cfg.DB)
	if err != nil {
		return err
	}
	utils.Info("database initialized")
	tmp, _ := db.DB()
	defer tmp.Close()
	// Init todo repo.
	var todo todoRepo.Repository
	todo = todoDB.New(db)
	utils.Info("repository payment initialized")

	// Init service.
	service := service.New(
		todo,
	)

	utils.Info("service initialized")
	// Init web server.
	httpServer := http.New(http.Config{
		Port:            cfg.App.Port,
		ReadTimeout:     cfg.App.ReadTimeout,
		WriteTimeout:    cfg.App.WriteTimeout,
		GracefulTimeout: cfg.App.GracefulTimeout,
	})
	utils.Info("http server initialized")
	r := httpServer.Router()
	// Register api route.
	httpAPI.New(service, cfg.App.InternalKey, cfg.App.Env).Register(r)
	utils.Info("http route api initialized")
	// Run web server.
	httpServer.Run()
	utils.Info("http server listening at :%s", cfg.App.Port)
	return nil	
}