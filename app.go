package main

import (
	"github.com/dkowalsky/brieefly/config"
	"github.com/dkowalsky/brieefly/db"
	"github.com/dkowalsky/brieefly/log"
	"github.com/dkowalsky/brieefly/net"
)

func main() {
	c, err := config.NewConfig(config.Local)
	if err != nil {
		log.Error(err)
		return
	}

	db, err := db.Connect(c)
	if err != nil {
		log.Error(err)
		return
	}

	router := net.BrieeflyRouter(db, c)

	router.Run()
}
