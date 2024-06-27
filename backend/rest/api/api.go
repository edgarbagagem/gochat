package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/edgarbagagem/gochat/config"
	"github.com/edgarbagagem/gochat/services/user"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

// Constructor
func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

// Health check handler
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (s *APIServer) Run() {
	log.Printf("API Server is starting at %s", s.addr)

	router := mux.NewRouter()
	router.HandleFunc("/health", healthCheckHandler).Methods("GET")
	router.HandleFunc("/health/", healthCheckHandler).Methods("GET")
	router.HandleFunc("/api/health", healthCheckHandler).Methods("GET")
	router.HandleFunc("/api/health/", healthCheckHandler).Methods("GET")
	subrouter := router.PathPrefix("/api/v1/").Subrouter()

	//register services
	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	// Configure CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", config.Envs.FrontendURL},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// Use the CORS middleware
	handler := c.Handler(router)

	log.Fatal(http.ListenAndServe(s.addr, handler))
}
