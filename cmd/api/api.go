package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jpecheverryp/tiago-ecom-api/service/user"
)

type APIServer struct {
    addr string
    db *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
    return &APIServer{
        addr: addr,
        db: db,
    }
}

func (s *APIServer) Run() error {
    router := mux.NewRouter()
    subRouter := router.PathPrefix("/api/v1").Subrouter()

    userHandler := user.NewHandler()
    userHandler.RegisterRoutes(subRouter)


    log.Println("Listening on ", s.addr)

    return http.ListenAndServe(s.addr, router)
}
