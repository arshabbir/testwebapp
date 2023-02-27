package app

import (
	"fmt"
	"log"
	"testwebservermod/api"
	"testwebservermod/config"
)

type app struct {
	server api.Server
	conf   *config.Config
}

type App interface {
	Start() error
}

func NewApp(conf *config.Config) App {
	return &app{server: api.NewServer(conf)}
}

func (a *app) Start() error {
	fmt.Println("Starting the app.....")
	mux, err := a.server.AddRoutes()
	if err != nil {
		log.Println("error registering the routes")
		return err
	}

	return a.server.Start(mux)
}
