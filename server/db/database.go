package db

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/dkowalsky/brieefly/config"
	"github.com/dkowalsky/brieefly/err"
	"github.com/dkowalsky/brieefly/log"

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
func Connect(config *config.Config) (*DB, *err.Error) {
	connectionString := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true",
		config.Database.User,
		config.Database.Password,
		config.Database.Address,
		config.Database.Port,
		config.Database.Name)
	db, sqlErr := sql.Open("mysql", connectionString)
	if sqlErr != nil {
		return nil, err.New(sqlErr, err.ErrConnectionFailure, nil)
	}
	sqlErr = db.Ping()
	if sqlErr != nil {
		return nil, err.New(sqlErr, err.ErrConnectionFailure, nil)
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

// ParseSlug -
func ParseSlug(slug string) string {
	return strings.Replace(strings.ToLower(slug), " ", "-", -1)
}

// InsertStmt -
func InsertStmt(tx *sql.Tx, args interface{}, table string) FinishedInsert {
	insertMap := parseStruct(args)
	preparedInsert := parseInsertMap(insertMap)
	insert := "INSERT INTO"
	values := "VALUES"
	stmt := fmt.Sprintf("%s %s %s %s %s", insert, table, preparedInsert.cols, values, preparedInsert.args)
	sqlStmt, err := tx.Prepare(stmt)
	if err != nil {
		panic(err)
	}

	fi := FinishedInsert{sqlStmt, preparedInsert.vals}

	return fi
}

// UpdateStmt -
func UpdateStmt(tx *sql.Tx, args interface{}, table string, condition *string) FinishedUpdate {
	updateMap := parseStruct(args)
	preparedUpdate := parseUpdateMap(updateMap)
	update := "UPDATE"
	set := "SET"
	where := "WHERE"
	stmt := "" 
	if condition != nil {
		stmt = fmt.Sprintf("%s %s %s %s %s %s", update, table, set, preparedUpdate.sets, where, *condition)
	} else {
		stmt = fmt.Sprintf("%s %s %s %s", update, table, set, preparedUpdate.sets)
	}
	log.Debug(stmt)
	sqlStmt, err := tx.Prepare(stmt)
	if err != nil {
		panic(err)
	}

	fi := FinishedUpdate{sqlStmt, preparedUpdate.vals}

	return fi
}

// FinishedUpdate -
type FinishedUpdate struct {
	*sql.Stmt
	Args []interface{}
}

type preparedUpdate struct {
	sets string
	vals []interface{}
}

func parseUpdateMap(arg map[string]interface{}) preparedUpdate {
	upVals := preparedUpdate{sets: "", vals: []interface{}{}}
	size := len(arg)
	count := 0
	for k, v := range arg {
		upVals.sets += fmt.Sprintf("%v = ", k)
		if count == size-1 {
			upVals.sets += "?"
		} else {
			upVals.sets += "?, "
		}
		upVals.vals = append(upVals.vals, v)
		count++
	}
	return upVals
}

// FinishedInsert -
type FinishedInsert struct {
	*sql.Stmt
	Args []interface{}
}

type preparedInsert struct {
	cols string
	args string
	vals []interface{}
}

func parseInsertMap(arg map[string]interface{}) preparedInsert {
	insVals := preparedInsert{cols: "", args: "", vals: []interface{}{}}
	size := len(arg)
	count := 0
	for k, v := range arg {
		if count == 0 {
			insVals.cols += fmt.Sprintf("(%v, ", k)
			insVals.args += "(?, "
		} else if count == size-1 {
			insVals.cols += fmt.Sprintf("%v)", k)
			insVals.args += "?) "
		} else {
			insVals.cols += fmt.Sprintf("%v, ", k)
			insVals.args += "?, "
		}
		insVals.vals = append(insVals.vals, v)
		count++
	}
	return insVals
}

func parseStruct(arg interface{}) map[string]interface{} {
	v := reflect.ValueOf(arg)
	if v.Kind() != reflect.Struct {
		v = reflect.ValueOf(arg).Elem()
	}
	t := reflect.TypeOf(arg)
	if t.Kind() != reflect.Struct {
		t = reflect.TypeOf(arg).Elem()
	}
	res := map[string]interface{}{}

	for i := 0; i < v.NumField(); i++ {
		val := v.Field(i)
		field := t.Field(i)
		tag := field.Tag.Get(tagKey)
		if tag == "-" {
			fmt.Println("skipping")
			continue
		}
		switch val.Kind() {
		case reflect.Struct:
			value := val.Interface()
			switch value.(type) {
			case NullString:
				ns := value.(NullString)
				if ns.Valid {
					res[tag] = ns.String
				}
			case NullInt64:
				ni := value.(NullInt64)
				if ni.Valid {
					res[tag] = ni.Int64
				}
			case NullFloat64:
				nf := value.(NullFloat64)
				if nf.Valid {
					res[tag] = nf.Float64
				}
			case NullTime:
				nt := value.(NullTime)
				if nt.Valid {
					res[tag] = nt.Time
				}
			case NullBool:
				nb := value.(NullBool)
				if nb.Valid {
					res[tag] = nb.Bool
				}
			case time.Time:
				t := value.(time.Time)
				res[tag] = t
			default:
				nested := parseStruct(val.Interface())
				for k, v := range nested {
					res[k] = v
				}
			}
			//case reflect.Slice:
			// value := val.Interface()
			// switch value.(type) {
			// case []struct{}:
			// 	arr := value.([]struct{})
			// 	for _, v := range arr {
			// 		nested := parseStruct(v)
			// 		for k, v := range nested {
			// 			res[k] = v
			// 		}
			// 	}
			// }
		default:
			fmt.Printf("type: %s, kind: %s value: %s \n", field.Type, val.Kind(), val.String())
			res[tag] = val.Interface()
		}
	}
	return res
}
