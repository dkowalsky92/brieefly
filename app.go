package main // import "github.com/dkowalsky/brieefly"

import (
	"fmt"

	"github.com/brieefly/config"
	"github.com/brieefly/db"
	"github.com/brieefly/log"
	"github.com/brieefly/net"
)

func main() {
	fmt.Println("Configuring...")
	c, err := config.NewConfig(config.Local)
	if err != nil {
		log.Error(err)
		return
	}
	fmt.Println("Configuration successful.")

	fmt.Println("Connecting to database...")
	db, err := db.Connect(c)
	if err != nil {
		log.Error(err)
		return
	}
	fmt.Println("Connected.")

	router := net.BrieeflyRouter(db, c)

	fmt.Println("Server engine starting...")
	router.Run()
}
