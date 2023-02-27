package main

import (
	"log"
	"testwebservermod/app"
	"testwebservermod/config"
)

func main() {

	conf := &config.Config{Addr: ":8080"}
	webapp := app.NewApp(conf)

	if err := webapp.Start(); err != nil {
		log.Fatal("error starting the webapp...")
	}

}
