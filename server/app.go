package main

import (
	"fmt"

	"github.com/brieefly/server/config"
	"github.com/brieefly/server/db"
	"github.com/brieefly/server/log"
	"github.com/brieefly/server/net"
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
