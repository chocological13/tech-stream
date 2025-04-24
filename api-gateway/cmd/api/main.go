package main

import (
	"context"
	"errors"
	"github.com/chocological13/tech-stream/api-gateway/internal/client"
	"github.com/chocological13/tech-stream/api-gateway/internal/config"
	"github.com/chocological13/tech-stream/api-gateway/internal/handlers"
	"github.com/chocological13/tech-stream/api-gateway/internal/middleware"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// TODO : initialize clients
	userClient, err := client.NewUserClient(cfg.UserServiceAddress)
	if err != nil {
		log.Fatalf("Error creating user client: %v", err)
	}

	// TODO : initialize handlers
	authHandler := handlers.NewAuthHandler(userClient)

	// Setup router
	r := gin.Default()

	// Add middleware
	r.Use(middleware.Logger())
	r.Use(middleware.RequestID())
	r.Use(middleware.CORS())

	// Healthcheck
	r.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	// TODO : setup routes
	api := r.Group("/api/v1")
	{
		// Auth routes
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
		}

		// Protected routes
		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			// TODO : add protected routes
		}
	}

	// Create server
	srv := &http.Server{
		Addr:    cfg.Address,
		Handler: r,
	}

	// Start server in goroutine
	go func() {
		log.Infof("Starting server on %s", cfg.Address)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Infof("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Error shutting down server: %v", err)
	}

	log.Infof("Server stopped")
}
