package main

import (
	"context"
	"loans/bootstrap"
	"loans/delivery/routes"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	log.Println("Starting server...")
	err := bootstrap.InitEnv()
	if err != nil {
		log.Fatalf("Error initializing environment: %v", err)
	}

	log.Println("Connecting to MongoDB...")
	uri, err := bootstrap.GetEnv("MONGODB_URI")
	if err != nil {
		log.Fatalf("Error getting environment variable MONGODB_URI: %v", err)
	}

	client, err := bootstrap.ConnectToMongoDB(uri)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	log.Println("Starting HTTP server...")
	router := routes.InitRoutes(client)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Run server in a goroutine
	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Println("Server is running on port 8080")

	// Graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()
	log.Println("Shutting down server...")

	// Create a context with a timeout for the graceful shutdown
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Close database connection
	if err := bootstrap.DisconnectFromMongoDB(client); err != nil {
		log.Fatalf("Error disconnecting from MongoDB: %s", err)
	} else {
		log.Println("Database connection closed")
	}

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server forced to shutdown: %s", err)
	}

	log.Println("Server exiting")
}
