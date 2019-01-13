package db

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/brieefly/config"
	"github.com/brieefly/err"
	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// error constants
const (
	ErrNotFound err.ErrorType = iota
	ErrAlreadyExists
)

// DB - sql.DB wrapper
type DB struct {
	*sql.DB
}

// Connect - connect to a database and return it
func Connect(config *config.Config) (*DB, error) {
	connectionString := fmt.Sprintf("%s:%s@/%s?parseTime=true", config.Database.User, config.Database.Password, config.Database.Name)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

// Disconnect - disconnect from specified database
func Disconnect(db *DB) error {
	err := db.Close()
	return err
}

// HandleError - returns an appropriate error for
func (db *DB) HandleError(_err error) *err.Error {
	if _err != nil {
		switch _err {
		case sql.ErrNoRows:
			return err.New(_err, err.ErrEmptyResult, map[string]interface{}{})
		default:
			return err.New(_err, err.ErrMalformedQuery, map[string]interface{}{})
		}
	}

	return nil
}

// HandleTypedError - returns an appropriate error for
func (db *DB) HandleTypedError(_err error, t err.ErrorType) *err.Error {
	switch t {
	case ErrNotFound:
		return err.New(errors.New("specified resource does not exist"), err.ErrEmptyResult, map[string]interface{}{})
	case ErrAlreadyExists:
		return err.New(errors.New("specified resource already exists and cannot be replaced"), err.ErrConflict, map[string]interface{}{})
	default:
		return err.New(_err, err.ErrEmptyResult, map[string]interface{}{})
	}
}
