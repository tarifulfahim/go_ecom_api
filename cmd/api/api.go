package api

import (
	"go_ecom_api/cmd/services/user"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
)

type APIServer struct {
	addr string
	db   *pgx.Conn
}

func NewServerAPI(addr string, db *pgx.Conn) *APIServer {
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
