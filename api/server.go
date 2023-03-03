package api

import (
	"fmt"
	"log"
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

	s.MiddleWare(mux)
	mux.Handle("/ping", s.MiddleWare(http.HandlerFunc(s.handlePing)))
	mux.Handle("/home", s.MiddleWare(http.HandlerFunc(s.handleHome)))
	// mux.Handle("/create", s.MiddleWare(http.HandlerFunc(s.handleCreate)))
	s.MiddleWare(mux)

	return mux, nil

}

func (s *server) MiddleWare(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Middleware invoked : before handler")
		defer log.Println("Middleinvoked : after handler")
		h.ServeHTTP(w, r)
	})

}
func (s *server) Start(mux http.Handler) error {
	fmt.Println("Starting the Server....")
	return http.ListenAndServe(s.conf.Addr, mux)

}
