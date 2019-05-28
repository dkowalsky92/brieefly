package main

import (
	// "os"
	// "bufio"
	// "errors"

	"github.com/dkowalsky/brieefly/config"
	"github.com/dkowalsky/brieefly/db"
	"github.com/dkowalsky/brieefly/log"
	"github.com/dkowalsky/brieefly/net"
	"github.com/dkowalsky/brieefly/retry"
	"github.com/dkowalsky/brieefly/err"
)

func main() {
	go func() {
		retry.PerformInfinite(retry.DefaultOptions(), func() *err.Error {
			log.Info("Configuring...")
			c, cErr := config.NewConfig(config.Local)
			if cErr != nil {
				return cErr
			}
			log.Info("Configuration successful.")
			
			log.Info("Connecting to database...")
			db, dbErr := db.Connect(c)
			if dbErr != nil {
				return dbErr
			}
			log.Info("Connected.")

			router := net.NewRouter(db, c)

			
			log.Info("Server is running.")
			log.Info("Accepting standard input -> ")
			rtErr := router.Run()
			if rtErr != nil {
				return rtErr
			}

			return nil
		})
	}()

	select {}
	// scanner := bufio.NewScanner(os.Stdin)
	// for scanner.Scan() {
	// 	log.Info("->")
	// }
}
