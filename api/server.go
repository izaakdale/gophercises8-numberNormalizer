package api

import (
	"log"
	"net/http"

	db "github.com/izaakdale/phoneNormalizer/db/sqlc"
)

type Server struct {
	store  db.Store
	router *http.ServeMux
}

func NewServer(store db.Store) (*Server, error) {

	server := &Server{
		store: store,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/numbers", server.PostNumber)
	mux.HandleFunc("/normalizednumbers", server.GetNormalizedNumbers)

	server.router = mux

	return server, nil
}

func (server *Server) Start() {
	log.Fatal(http.ListenAndServe(":8080", server.router))
}
