package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"gateway/internal/config"
	transporthttp "gateway/internal/transport/http"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load(".env")

	cfg, err := config.ReadEnv()
	if err != nil {
		log.Fatal("configuration error:", err)
	}

	handler, err := transporthttp.NewHandler(cfg)
	if err != nil {
		log.Fatal("handlers setup error:", err)
	}

	server := &http.Server{
		Addr:         cfg.ListenAddr(),
		Handler:      handler,
		ReadTimeout:  cfg.HTTP.ReadTimeout,
		WriteTimeout: cfg.HTTP.WriteTimeout,
		IdleTimeout:  cfg.HTTP.IdleTimeout,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		log.Println("gateway started on", cfg.ListenAddr())
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("server error:", err)
		}
	}()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.HTTP.ShutdownTimeout)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Println("shutdown error:", err)
	}

	log.Println("gateway stopped")
}
