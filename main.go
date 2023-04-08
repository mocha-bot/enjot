package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/mocha-bot/enjot/config"
	workspaceUsecase "github.com/mocha-bot/enjot/core/module"
	httpInternal "github.com/mocha-bot/enjot/handler/http"
	"github.com/mocha-bot/enjot/pkg/logger"
	middlewareInternal "github.com/mocha-bot/enjot/pkg/middleware"
	workspaceRepository "github.com/mocha-bot/enjot/repository"
)

func main() {
	cfg := config.Get()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

	db, err := gorm.Open(mysql.Open(cfg.Database.Mysql.DSN()), &gorm.Config{})

	if err != nil {
		logger.Fatal().
			Err(err).
			Caller().
			Str("context", "gorm.open.connection").
			Msg("failed open connection to database")
	}

	r := chi.NewRouter()

	// initialize repository
	workspaceRepo := workspaceRepository.NewWorkspaceRepository(db)

	// initialize usecase
	workspaceUC := workspaceUsecase.NewWorkspaceUsecase(workspaceRepo)

	r.Use(middlewareInternal.Logger())
	r.Use(middleware.Recoverer)

	// health check handler
	r.Get("/health", httpInternal.HealthCheck)

	httpInternal.NewDummyHandler(r)

	r.Route("/api", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(middlewareInternal.Authorization([]byte(cfg.TokenKeySecret)))
			httpInternal.NewWorkspaceHandler(r, workspaceUC)
		})
	})

	srv := &http.Server{
		Addr:         cfg.Server.Address(),
		WriteTimeout: cfg.Server.WriteTimeout,
		ReadTimeout:  cfg.Server.ReadTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
		Handler:      r,
	}

	go func() {
		logger.Info().
			Msgf("server running at %s", srv.Addr)

		err := srv.ListenAndServe()

		if err != nil && err != http.ErrServerClosed {
			logger.Error().
				Err(err).
				Str("context", "listen.and.serve").
				Send()
		}
	}()

	<-c

	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), cfg.Server.WaitGracefulTimeout)
	defer cancel()

	err = srv.Shutdown(ctxWithTimeout)

	if err != nil {
		logger.Error().
			Err(err).
			Str("context", "shutdown").
			Send()
	}

	logger.Info().Msg("shutting down")
}
