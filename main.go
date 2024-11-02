package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/odrling/zitadel-karaberus/server/exampleop"
	"github.com/odrling/zitadel-karaberus/server/storage"
)

func main() {
	addr, found := os.LookupEnv("LISTEN_ADDR")
	if !found {
		addr = "127.0.0.1"
	}
	port, found := os.LookupEnv("LISTEN_PORT")
	if !found {
		port = "9998"
	}
	issuer := fmt.Sprintf("http://%s:%s/", addr, port)

	// the OpenIDProvider interface needs a Storage interface handling various checks and state manipulations
	// this might be the layer for accessing your database
	// in this example it will be handled in-memory
	storage := storage.NewStorage(storage.NewUserStore(issuer))

	logger := slog.New(
		slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelDebug,
		}),
	)
	router := exampleop.SetupServer(issuer, storage, logger, false)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	logger.Info("server listening, press ctrl+c to stop", "addr", fmt.Sprintf("http://localhost:%s/", port))
	err := server.ListenAndServe()
	if err != http.ErrServerClosed {
		logger.Error("server terminated", "error", err)
		os.Exit(1)
	}
}
