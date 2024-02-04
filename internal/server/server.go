package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"server-template/internal/database"
)

type Server struct {
	port       int
	db         database.Service
	httpServer *http.Server
}

// Option is a function that configures the server.
type Option func(*Server)

func NewServer(options ...Option) *Server {
	s := &Server{
		port: 8080, // Default port
		db:   database.New(),
	}

	// Override defaults with any specified options
	for _, option := range options {
		option(s)
	}

	s.initHTTPServer()
	return s
}

// WithPort is an Option to configure the server port.
func WithPort(port int) Option {
	return func(s *Server) {
		s.port = port
	}
}

// WithDatabaseService is an Option to configure the database service.
func WithDatabaseService(db database.Service) Option {
	return func(s *Server) {
		s.db = db
	}
}

// initHTTPServer initializes the *http.Server with the Server's configuration.
func (s *Server) initHTTPServer() {
	s.httpServer = &http.Server{
		Addr:         fmt.Sprintf(":%d", s.port),
		Handler:      s.RegisterRoutes(), // Implement this method to setup routes
		IdleTimeout:  1 * time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
}

// Start runs the server, logging any errors encountered.
func (s *Server) Start() {
	log.Printf("Server starting on port %d\n", s.port)
	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed to start: %v", err)
	}
}

// loadPortFromEnv tries to load the server port from an environment variable, falling back to a default if not found or invalid.
func loadPortFromEnv() int {
	portStr, exists := os.LookupEnv("PORT")
	if !exists {
		return 8080 // Return default port if not set
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Printf("Warning: Invalid PORT environment variable '%s', falling back to default port 8080.\n", portStr)
		return 8080
	}

	return port
}
