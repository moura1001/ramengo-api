package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/moura1001/ramengo-api/src/handlers"
	utilapp "github.com/moura1001/ramengo-api/src/util/app"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("GET /broths", handlers.WithRequiredHeaders(handlers.HandleBrothList))
	router.HandleFunc("GET /proteins", handlers.WithRequiredHeaders(handlers.HandleProteinList))
	router.HandleFunc("POST /orders", handlers.WithRequiredHeaders(handlers.HandleOrderNew))
	router.HandleFunc("GET /healthcheck", func(w http.ResponseWriter, r *http.Request) {
		handlers.WriteJSON(w, http.StatusOK, map[string]string{
			"status":  "ok",
			"version": utilapp.Version,
		})
	})

	port := utilapp.GetEnv("PORT", "8080")
	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	slog.Info("Starting...")
	go func() {
		if err := server.ListenAndServe(); err != nil && http.ErrServerClosed != err {
			panic(err)
		}
	}()
	slog.Info("Server listening on port " + server.Addr)

	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGTERM, os.Interrupt, syscall.SIGINT)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	slog.Info("Stopping...")

	if err := server.Shutdown(ctx); err != nil {
		panic(err)
	}
	slog.Info("Server stopped")
}
