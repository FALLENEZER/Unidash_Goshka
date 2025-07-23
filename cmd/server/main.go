package main

import (
	"context"
	"github.com/fallenezer/Unidash_Goshka/internal/config"
	"github.com/fallenezer/Unidash_Goshka/internal/handler"
	"github.com/go-chi/chi/v5"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
)

func handlerMain() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := Run(ctx); err != nil {
		return err
	}

	return nil
}

func Run(ctx context.Context) error {
	cfg := config.NewConfig()

	router := chi.NewRouter()

	router.Handle("/swagger/*", http.StripPrefix("/swagger/", http.FileServer(http.Dir("web/swagger"))))
	handler.RegisterRoutes(router, handler.Dependencies{
		AssetsFS: http.Dir(cfg.AssetsDir),
	})

	server := http.Server{
		Addr:    cfg.ServerAddr,
		Handler: router,
	}

	go func() {
		<-ctx.Done()
		slog.Info("shutting down server")
		err := server.Shutdown(ctx)
		if err != nil {
			return
		}
	}()

	slog.Info("starting server", slog.String("addr", cfg.ServerAddr))
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func main() {
	if err := handlerMain(); err != nil {
		log.Fatal(err)
	}
}
