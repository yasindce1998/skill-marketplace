package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yasindce1998/skill-marketplace/api/routes"
	"github.com/yasindce1998/skill-marketplace/config"
)

var (
	mode = flag.String("mode", "server", "Application mode: server or cli")
)

func main() {
	flag.Parse()


	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Fatal("Failed to load configuration:", err)
	}

	// Initialize database connection
	database, err := db.InitDB(cfg.Database)
	if err != nil {
		logger.Fatal("Failed to connect to database:", err)
	}
	defer database.Close()

	switch *mode {
	case "server":
		runServer(cfg, logger, database)
	case "cli":
		runCLI(database, logger)
	default:
		logger.Fatal("Invalid mode. Use 'server' or 'cli'")
	}
}

func runServer(cfg *config.Config, logger *logger.Logger, db *db.Database) {
	// Initialize router
	router := routes.NewRouter(db, logger)

	// Create server
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Server.Port),
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Server run context
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, _ := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				logger.Fatal("Graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			logger.Fatal(err)
		}
		serverStopCtx()
	}()

	// Start server
	logger.Info(fmt.Sprintf("Starting server on port %d", cfg.Server.Port))
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logger.Fatal(err)
	}

	// Wait for server context to be stopped
	<-serverCtx.Done()
}

func runCLI(db *db.Database, logger *logger.Logger) {
	cli := admin.NewCLI(db, logger)
	if err := cli.Run(); err != nil {
		logger.Fatal("CLI execution failed:", err)
	}
}