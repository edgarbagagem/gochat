package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/edgarbagagem/gochat/services/user"
	"github.com/gorilla/mux"
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

	log.Fatal(http.ListenAndServe(s.addr, subrouter))
}
