package httpServer

import (
	"context"
	"log"
	"net/http"
	"time"

	"microservices/api-gateway/http-server/handlers"
	"microservices/api-gateway/http-server/routes"
	"microservices/pkg/proto-gen/authpb"
	"microservices/pkg/proto-gen/bookingpb"
	"microservices/pkg/proto-gen/userpb"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

const (
	defaultHTTPPort     = ":8000"
	defaultReadTimeout  = 5 * time.Second
	defaultWriteTimeout = 5 * time.Second
	defaultIdleTimeout  = 30 * time.Second
)

type Server struct {
	httpServer     *http.Server
	UserService    userpb.UserServiceClient
	BookingService bookingpb.BookingServiceClient
	AuthService    authpb.AuthServiceClient
}

type Config struct {
	Port                 string
	ReadTimeout          time.Duration
	WriteTimeout         time.Duration
	IdleTimeout          time.Duration
	UserServiceClient    userpb.UserServiceClient
	BookingServiceClient bookingpb.BookingServiceClient
	AuthServiceClient    authpb.AuthServiceClient
}

func New(cfg Config) *Server {
	applyDefaults(&cfg)

	router := chi.NewRouter()
	registerMiddleware(router)

	bookingHandler := &handlers.BookingHandler{Service: cfg.BookingServiceClient}
	userHandler := &handlers.UserHandler{Service: cfg.UserServiceClient}
	authHandler := &handlers.AuthHandler{Service: cfg.AuthServiceClient}

	router.Route("/api", func(router chi.Router) {
		router.Mount("/users", routes.UserRouter(userHandler))
		router.Mount("/bookings", routes.BookingRouter(bookingHandler))
		router.Mount("/auth", routes.AuthRouter(authHandler))
	})

	// Fallback 404 handler
	router.NotFound(func(w http.ResponseWriter, router *http.Request) {
		http.Error(w, "Route not found", http.StatusNotFound)
	})

	server := &http.Server{
		Addr:         cfg.Port,
		Handler:      router,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	return &Server{
		httpServer:     server,
		UserService:    cfg.UserServiceClient,
		BookingService: cfg.BookingServiceClient,
		AuthService:    cfg.AuthServiceClient,
	}
}

func registerMiddleware(router chi.Router) {
	router.Use(middleware.Logger)
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
}

func applyDefaults(cfg *Config) {
	if cfg.Port == "" {
		cfg.Port = defaultHTTPPort
	}
	if cfg.ReadTimeout == 0 {
		cfg.ReadTimeout = defaultReadTimeout
	}
	if cfg.WriteTimeout == 0 {
		cfg.WriteTimeout = defaultWriteTimeout
	}
	if cfg.IdleTimeout == 0 {
		cfg.IdleTimeout = defaultIdleTimeout
	}
}

// Start begins listening on the configured port
func (s *Server) Start() error {
	log.Printf("Starting HTTP server on %s\n", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

// Shutdown gracefully shuts down the server
func (s *Server) Shutdown(timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	log.Println("Shutting down HTTP server...")
	return s.httpServer.Shutdown(ctx)
}
