package main

import (
	"fmt"

	"github.com/dkowalsky/brieefly/config"
	"github.com/dkowalsky/brieefly/db"
	"github.com/dkowalsky/brieefly/log"
	"github.com/dkowalsky/brieefly/net"
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

	router := net.NewRouter(db, c)

	fmt.Println("Server is running.")
	router.Run()
}
