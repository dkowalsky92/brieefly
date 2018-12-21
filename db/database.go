package db

import (
	"database/sql"
	"fmt"

	"github.com/dkowalsky/brieefly/config"
	"github.com/dkowalsky/brieefly/log"
	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// DB - sql.DB wrapper
type DB struct {
	*sql.DB
}

// Connect - connect to a database and return it
func Connect(config *config.Config) (*DB, error) {
	connectionString := fmt.Sprintf("%s:%s@/%s", config.Database.User, config.Database.Password, config.Database.Name)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Error(err, "that connection error")
		return nil, err
	}

	return &DB{db}, nil
}

// Disconnect - disconnect from specified database
func Disconnect(db *DB) error {
	err := db.Close()
	return err
}
