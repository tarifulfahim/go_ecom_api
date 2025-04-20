package api

import (
	"database/sql"
	"go_ecom_api/cmd/services/user"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewServerAPI(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {

	router := mux.NewRouter()

	subRouter := router.PathPrefix("/api/v1").Subrouter()

	userHandler := user.NewHandler()

	userHandler.RegisterRoutes(subRouter)

	log.Println("server is running on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
