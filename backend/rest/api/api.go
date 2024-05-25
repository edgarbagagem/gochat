package api

import (
	"log"
	"net/http"

	db "github.com/edgarbagagem/gochat/db"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr  string
	store db.Store
}

// Constructor
func NewAPIServer(addr string, store db.Store) *APIServer {
	return &APIServer{
		addr:  addr,
		store: store,
	}
}

func (s *APIServer) Run() {
	log.Printf("API Server is starting at %s", s.addr)

	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1/").Subrouter()

	//register services

	log.Fatal(http.ListenAndServe(s.addr, subrouter))
}
