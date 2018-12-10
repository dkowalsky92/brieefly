package database

import (
	"database/sql"
	"fmt"

	"github.com/dkowalsky/brieefly/config"
	"github.com/dkowalsky/brieefly/log"
	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// Database - a database wrapper allowing connection to brieefly's MySQL database
type Database struct {
	*sql.DB
}

// Connect - connect to a database and return it
func Connect(config *config.Config) (*Database, error) {
	connectionString := fmt.Sprintf("%s:%s@/%s", config.Database.User, config.Database.Password, config.Database.Name)
	database, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Error(err, "that connection error")
		return nil, err
	}

	return &Database{database}, nil
}

// Disconnect - disconnect from specified database
func Disconnect(db *Database) error {
	err := db.Close()
	return err
}

// NullString - null string
func NullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

// NullInt - null int
func NullInt(i int64) sql.NullInt64 {
	if i == -1 {
		return sql.NullInt64{}
	}
	return sql.NullInt64{
		Int64: i,
		Valid: true,
	}
}
