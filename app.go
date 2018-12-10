package main

import (
	"github.com/dkowalsky/brieefly/config"
	"github.com/dkowalsky/brieefly/database"
	"github.com/dkowalsky/brieefly/log"
	"github.com/dkowalsky/brieefly/net"
)

func main() {

	c, err := config.NewConfig(config.Local)
	if err != nil {
		log.Error(err)
		return
	}
	log.Debug(c.Server.Certificate, c.Server.Key)

	db, err := database.Connect(c)
	if err != nil {
		log.Error(err)
		return
	}

	router := net.NewRouter(db, c)

	router.Run()
}
