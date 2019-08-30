package main

import (
	"bufio"
	"os"

	"github.com/dkowalsky/brieefly/config"
	"github.com/dkowalsky/brieefly/db"
	"github.com/dkowalsky/brieefly/err"
	"github.com/dkowalsky/brieefly/log"
	"github.com/dkowalsky/brieefly/net"
	"github.com/dkowalsky/brieefly/retry"
)

// func shutdown() {
// 	s := make(chan os.Signal, 1)
// 	signal.Notify(s, os.Interrupt)
// 	signal.Notify(s, syscall.SIGTERM)
// 	go func() {
// 		<-s
// 		fmt.Println("Sutting down gracefully.")
// 		// clean up here
// 		os.Exit(0)
// 	}()
// }

const (
	runtimeEnvKey string = "BRIEEFLY_ENVIRONMENT"
)

func envFromArgs(arg string) config.Environment {
	switch arg {
	case "dev", "development":
		return config.Development
	case "prod", "production":
		return config.Production
	case "local":
		return config.Local
	default:
		return config.Local
	}
}

func main() {
	go func() {
		retry.PerformInfinite(retry.DefaultOptions(), func() *err.Error {
			log.Info("Configuring...")
			env := envFromArgs(os.Getenv(runtimeEnvKey))
			c, cErr := config.NewConfig(env)
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
			log.Info("Accepting standard input: ")
			rtErr := router.Run()
			if rtErr != nil {
				return rtErr
			}

			return nil
		})
	}()

	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n')
		log.Info(input)
		switch input {
		case "exit":
			log.Warning("Shutting down...")
			os.Exit(0)
		default:
		}
	}
}
