package main

import (
	// "os"
	// "bufio"
	// "errors"

	"github.com/dkowalsky/brieefly/config"
	"github.com/dkowalsky/brieefly/db"
	"github.com/dkowalsky/brieefly/log"
	"github.com/dkowalsky/brieefly/net"
	//"github.com/dkowalsky/brieefly/retry"
	//"github.com/dkowalsky/brieefly/err"
)

func main() {
	//retry.PerformInfinite(retry.DefaultOptions(), func() *err.Error {
		log.Info("Configuring...")
		c, cErr := config.NewConfig(config.Local)
		if cErr != nil {
			panic(cErr)
		//	return cErr
		}
		log.Info("Configuration successful.")
		
		log.Info("Connecting to database...")
		db, dbErr := db.Connect(c)
		if dbErr != nil {
			log.Error(dbErr)
			panic(dbErr)
			//return dbErr
		}
		log.Info("Connected.")

		router := net.NewRouter(db, c)

		
		log.Info("Server is running.")
		log.Info("Accepting standard input -> ")
		rtErr := router.Run()
		if rtErr != nil {
			panic(rtErr)
		//	return rtErr
		}

	//	return nil
	//})
	// scanner := bufio.NewScanner(os.Stdin)
	// for scanner.Scan() {
	// 	log.Info("->")
	// }
}
