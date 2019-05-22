package db

import (
	"database/sql"

	"github.com/dkowalsky/brieefly/err"
)

// TxFn - is a function that will be called with an initialized `sql.Tx` object
// that can be used for executing statements and queries against a database.
type TxFn func(*sql.Tx) *err.Error

// WithTransaction - creates a new transaction and handles rollback/commit based on the
// error object returned by the `TxFn`
func (db *DB) WithTransaction(fn TxFn) *err.Error {
	tx, errTx := db.Begin()

	if errTx != nil {
		return err.New(errTx, err.ErrTxBegin, map[string]interface{}{})
	}

	defer func() {
		if errTx != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	txErr := fn(tx)
	return txErr
}
