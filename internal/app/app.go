package app

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pizzament/rsc-test/internal/app/handlers/counter_handler"
	"github.com/pizzament/rsc-test/internal/app/handlers/stats_handler"
	"github.com/pizzament/rsc-test/internal/infra/config"
	"github.com/pizzament/rsc-test/internal/repository"
	"github.com/pizzament/rsc-test/internal/service"
)

type App struct {
	config *config.Config
	server http.Server
}

func NewApp(configPath string) (*App, error) {
	configImpl, err := config.LoadConfig(configPath)
	if err != nil {
		return nil, fmt.Errorf("config.LoadConfig: %w", err)
	}

	app := &App{
		config: configImpl,
	}

	app.server.Handler = bootstrapHandler(configImpl)

	return app, nil
}

func (app *App) ListenAndServe() error {
	address := fmt.Sprintf("%s:%s", app.config.Service.Host, app.config.Service.Port)

	l, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	return app.server.Serve(l)
}

func bootstrapHandler(config *config.Config) http.Handler {
	ctx := context.Background()

	// бд пул
	configPgx, err := pgxpool.ParseConfig(fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.DBName))
	if err != nil {
		log.Fatalln("unable to parse pgx config: ", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, configPgx)
	if err != nil {
		log.Fatalln("unable to create pgx pool: ", err)
	}

	// создание репозитория и сервиса
	repo := repository.NewRepository(pool)
	svc := service.NewService(repo)

	mx := http.NewServeMux()
	mx.Handle("GET /counter/{banner_id}", counter_handler.NewCounterHandler(svc))
	mx.Handle("POST /stats/{banner_id}", stats_handler.NewStatsHandler(svc))

	return mx
}
