package db

import (
	"database/sql"
	"encoding/json"

	"github.com/go-sql-driver/mysql"
)

// NullInt64 - an alias for sql.NullInt64 data type
type NullInt64 struct {
	sql.NullInt64
}

// MarshalJSON - NullInt64 json marshalling with javascript null handling
func (ni *NullInt64) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Int64)
}

// NullFloat64 - an alias for sql.NullFloat64 data type
type NullFloat64 struct {
	sql.NullFloat64
}

// MarshalJSON - NullFloat64 json marshalling with javascript null handling
func (nf *NullFloat64) MarshalJSON() ([]byte, error) {
	if !nf.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nf.Float64)
}

// NullString - an alias for sql.NullString data type
type NullString struct {
	sql.NullString
}

// MarshalJSON - NullString json marshalling with javascript null handling
func (ns *NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

// NullTime - an alias for sql.NullTime data type
type NullTime struct {
	mysql.NullTime
}

// MarshalJSON - NullTime json marshalling with javascript null handling
func (nt *NullTime) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nt.Time)
}
