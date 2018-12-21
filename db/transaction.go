package db

import "database/sql"

// TxFn - is a function that will be called with an initialized `sql.Tx` object
// that can be used for executing statements and queries against a database.
type TxFn func(*sql.Tx) error

// WithTransaction - creates a new transaction and handles rollback/commit based on the
// error object returned by the `TxFn`
func (db *DB) WithTransaction(fn TxFn) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	err = fn(tx)
	return err
}
