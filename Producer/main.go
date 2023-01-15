package main

import (
	"context"
	"demo-k8s-messagequeue/api"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Http server configuration point to http.Server
	server := &http.Server{
		Addr:         ":8080",
		Handler:      api.ApiV1(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	go func() {
		// Create a channel to receive OS signals
		sigs := make(chan os.Signal, 1)

		// Listen for specific signals
		signal.Notify(sigs, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		sig := <-sigs
		log.Println("received signal:", sig)
		// Perform cleanup here before exiting
		os.Exit(0)
	}()

	// Start http server and gracefully shutdown
	log.Println("Starting server on port 8080")
	if err := server.ListenAndServe(); err != nil {
		log.Println("Server shutdown")
	}
	// Gracefully shutdown the server, waiting max 5 seconds for current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")

}
