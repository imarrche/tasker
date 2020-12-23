package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

// Server is the REST API server serving frontends.
type Server struct {
	l      *log.Logger
	config *Config
	router *mux.Router
}

// NewServer creates a new Server instance.
func NewServer() *Server {
	l := log.New(os.Stdout, "", log.LstdFlags)
	c := NewConfig()     // Reading config fron environment.
	r := mux.NewRouter() // Initializing router.

	return &Server{l: l, config: c, router: r}
}

// Start starts the server.
func (s *Server) Start() {
	// Initializing HTTP server.
	server := &http.Server{
		Addr:    s.config.Addr,
		Handler: s.router,
	}

	// Setting up routes.
	s.router.HandleFunc("/", testHandler).Methods("GET") // Test endpoint.

	// Graceful shutdown setup.
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	// Starting the server.
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.l.Fatal("couldn't start the server")
		}
	}()
	s.l.Println("server started")

	<-done
	// Gracefully shutting down the server.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()
	if err := server.Shutdown(ctx); err != nil {
		s.l.Fatal("server couldn't gracefully shut down")
	}
	s.l.Println("server shutted down gracefully")
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Tasker"))
}
