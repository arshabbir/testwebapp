package api

import (
	"fmt"
	"net/http"
	"testwebservermod/config"
)

type server struct {
	conf *config.Config
}

type Server interface {
	Start(mux http.Handler) error
	AddRoutes() (http.Handler, error)
}

func NewServer(conf *config.Config) Server {
	return &server{conf: conf}
}

func (s *server) AddRoutes() (http.Handler, error) {

	mux := http.NewServeMux()

	//s.LogMiddleWare(mux)
	mux.Handle("/ping", s.LogMiddleWare(http.HandlerFunc(s.handlePing)))
	mux.Handle("/home", s.IPExtractMiddleWare(http.HandlerFunc(s.handleHome)))
	// mux.Handle("/create", s.MiddleWare(http.HandlerFunc(s.handleCreate)))
	s.LogMiddleWare(mux)

	return mux, nil

}

func (s *server) Start(mux http.Handler) error {
	fmt.Println("Starting the Server....")
	return http.ListenAndServe(s.conf.Addr, mux)

}
