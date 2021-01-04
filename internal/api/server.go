package api

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"

	"github.com/imarrche/tasker/internal/config"
	"github.com/imarrche/tasker/internal/service"
	"github.com/imarrche/tasker/internal/service/web"
	"github.com/imarrche/tasker/internal/store"
	"github.com/imarrche/tasker/internal/store/inmem"
)

// Server is the REST API server for Tasker.
type Server struct {
	l       *log.Logger
	config  *config.Config
	store   store.Store
	router  *mux.Router
	service service.Service
}

// NewServer creates a new Server instance.
func NewServer(l *log.Logger, c *config.Config, store store.Store) *Server {
	r := mux.NewRouter()
	service := web.NewService(store)

	return &Server{l: l, config: c, store: store, router: r, service: service}
}

// NewTestServer creates a new test Server instance.
func NewTestServer() *Server {
	l := log.New(os.Stdout, "", log.LstdFlags)
	c := config.New()
	store := inmem.TestStoreWithFixtures()
	r := mux.NewRouter()
	service := web.NewService(store)

	return &Server{l: l, config: c, store: store, router: r, service: service}
}

// Start starts the server.
func (s *Server) Start() error {
	// Initializing HTTP server.
	server := &http.Server{
		Addr:         s.config.Addr,
		Handler:      s.router,
		IdleTimeout:  60 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
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
		return errors.New("server couldn't gracefully shut down")
	}
	if err := s.store.Close(); err != nil {
		return errors.New("couldn't close the store")
	}

	s.l.Println("server shutted down gracefully")

	return nil
}

func (s *Server) configureRouter() {
	v1Router := s.router.PathPrefix("/api/v1").Subrouter()

	projects := v1Router.PathPrefix("/projects").Subrouter()
	projects.HandleFunc("", s.projectList()).Methods(http.MethodGet)
	projects.HandleFunc("", s.projectCreate()).Methods(http.MethodPost)
	projects.HandleFunc("/{project_id:[0-9]+}", s.projectDetail()).Methods(http.MethodGet)
	projects.HandleFunc("/{project_id:[0-9]+}", s.projectUpdate()).Methods(http.MethodPut)
	projects.HandleFunc("/{project_id:[0-9]+}", s.projectDelete()).Methods(http.MethodDelete)
	projects.HandleFunc("/{project_id:[0-9]+}/columns", s.columnList()).Methods(http.MethodGet)
	projects.HandleFunc("/{project_id:[0-9]+}/columns", s.columnCreate()).Methods(http.MethodPost)

	columns := v1Router.PathPrefix("/columns").Subrouter()
	columns.HandleFunc("/{column_id:[0-9]+}", s.columnDetail()).Methods(http.MethodGet)
	columns.HandleFunc("/{column_id:[0-9]+}", s.columnUpdate()).Methods(http.MethodPut)
	columns.HandleFunc("/{column_id:[0-9]+}/move", s.columnMove()).Methods(http.MethodPost)
	columns.HandleFunc("/{column_id:[0-9]+}", s.columnDelete()).Methods(http.MethodDelete)
	columns.HandleFunc("/{column_id:[0-9]+}/tasks", s.taskList()).Methods(http.MethodGet)
	columns.HandleFunc("/{column_id:[0-9]+}/tasks", s.taskCreate()).Methods(http.MethodPost)

	tasks := v1Router.PathPrefix("/tasks").Subrouter()
	tasks.HandleFunc("/{task_id:[0-9]+}", s.taskDetail()).Methods(http.MethodGet)
	tasks.HandleFunc("/{task_id:[0-9]+}", s.taskUpdate()).Methods(http.MethodPut)
	tasks.HandleFunc("/{task_id:[0-9]+}/movex", s.taskMoveX()).Methods(http.MethodPost)
	tasks.HandleFunc("/{task_id:[0-9]+}/movey", s.taskMoveY()).Methods(http.MethodPost)
	tasks.HandleFunc("/{task_id:[0-9]+}", s.taskDelete()).Methods(http.MethodDelete)
	tasks.HandleFunc("/{task_id:[0-9]+}/comments", s.commentList()).Methods(http.MethodGet)
	tasks.HandleFunc("/{task_id:[0-9]+}/comments", s.commentCreate()).Methods(http.MethodPost)

	comments := v1Router.PathPrefix("/comments").Subrouter()
	comments.HandleFunc("/{comment_id:[0-9]+}", s.commentDetail()).Methods(http.MethodGet)
	comments.HandleFunc("/{comment_id:[0-9]+}", s.commentUpdate()).Methods(http.MethodPut)
	comments.HandleFunc("/{comment_id:[0-9]+}", s.commentDelete()).Methods(http.MethodDelete)
}

func (s *Server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func (s *Server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	if code >= 500 {
		s.l.Printf("[SERVER ERROR]: %s\n", err.Error())
		err = nil // Do not show server error to users for security reasons.
	}
	if err != nil {
		s.respond(w, r, code, map[string]string{"error": err.Error()})
		return
	}
	s.respond(w, r, code, nil)
}
