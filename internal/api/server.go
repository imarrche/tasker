package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"

	"github.com/imarrche/tasker/internal/service"
	"github.com/imarrche/tasker/internal/service/web"
	"github.com/imarrche/tasker/internal/store"
)

// Server is the REST API server serving frontends.
type Server struct {
	l       *log.Logger
	config  *Config
	router  *mux.Router
	store   store.Store
	service service.Service
}

// NewServer creates a new Server instance.
func NewServer(store store.Store) *Server {
	l := log.New(os.Stdout, "", log.LstdFlags)
	c := NewConfig()     // Reading config fron environment.
	r := mux.NewRouter() // Initializing router.
	service := web.NewService(store)

	return &Server{l: l, config: c, router: r, store: store, service: service}
}

// Start starts the server.
func (s *Server) Start() {
	// Initializing HTTP server.
	server := &http.Server{
		Addr:    s.config.Addr,
		Handler: s.router,
	}

	// Setting up router.
	s.configureRouter()

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

func (s *Server) configureRouter() {
	s.router.HandleFunc("/projects", s.handleProjectList()).Methods("GET")
	s.router.HandleFunc("/projects", s.handleProjectCreate()).Methods("POST")
	s.router.HandleFunc("/projects/{id:[0-9]+}", s.handleProjectDetail()).Methods("GET")
	s.router.HandleFunc("/projects/{id:[0-9]+}", s.handleProjectUpdate()).Methods("PUT")
	s.router.HandleFunc("/projects/{id:[0-9]+}", s.handleProjectDelete()).Methods("DELETE")
}

func (s *Server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func (s *Server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	if err != nil {
		s.respond(w, r, code, map[string]string{"error": err.Error()})
	}
	s.respond(w, r, code, nil)
}
