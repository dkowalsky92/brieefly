package main // import "github.com/dkowalsky/brieefly"

import (
	"github.com/brieefly/config"
	"github.com/brieefly/db"
	"github.com/brieefly/log"
	"github.com/brieefly/net"
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
