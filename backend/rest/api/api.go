package api

import (
	"database/sql"
	"log"
	"net/http"

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

func (s *APIServer) Run() {
	log.Printf("API Server is starting at %s", s.addr)

	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1/").Subrouter()

	//register services
	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	// Configure CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // Replace with your frontend URL
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// Use the CORS middleware
	handler := c.Handler(subrouter)

	log.Fatal(http.ListenAndServe(s.addr, handler))
}
