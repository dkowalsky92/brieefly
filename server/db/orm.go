package db

import (
	"database/sql"
)

type keyword string

const (
	mInsert keyword = "INSERT"
	mDelete keyword = "UPDATE"
	mUpdate keyword = "DELETE"
	mSelect keyword = "SELECT"
	mInto   keyword = "INTO"
	mValues keyword = "VALUES"
	mFrom   keyword = "FROM"
)

const tagKey string = "orm"

type pattern map[string]interface{}

// Stmt -
type Stmt struct {
	//selects []pattern
	inserts []pattern
	//updates []pattern
	//deletes []pattern
	wheres []pattern
	ors    []pattern
	values []pattern
}

// Build -
func (st *Stmt) Build(tx *sql.Tx) (*sql.Stmt, error) {
	stmt := ""

	//if inserts != nil && vals[intoKey] != nil && vals[valuesKey] != nil {
	// cols := vals[columnsKey].([]string)
	// into := vals[intoKey].(keyword)
	// //vals := vals[valuesKey].([]string)

	// stmt = fmt.Sprintf("%v %v", method, into)
	// for i, v := range cols {
	// 	if i == len(cols)-1 {
	// 		stmt += fmt.Sprintf("%v ", v)
	// 	} else {
	// 		stmt += fmt.Sprintf("%v, ", v)
	// 	}
	// }
	// stmt += fmt.Sprintf("%v ")
	//}

	return tx.Prepare(stmt)
}

func build(p pattern) string {

	return ""
}

// Builder -
type Builder interface {
	InsertInto() *Stmt
	Insert() *Stmt
	Update() *Stmt
	Delete() *Stmt
	Select() *Stmt
	Into() *Stmt
	Columns(args interface{}) *Stmt
	Values(args interface{}) *Stmt
	From(args interface{}) *Stmt
	Where(condition interface{}, args interface{}) *Stmt
}

// Insert -
func (st *Stmt) Insert() *Stmt {
	// vals := st.pattern
	// vals[methodKey] = mInsert
	// st.pattern = vals
	return st
}

// Update -
func (st *Stmt) Update() *Stmt {
	// vals := st.pattern
	// vals[methodKey] = mUpdate
	// st.pattern = vals
	return st
}

// Delete -
func (st *Stmt) Delete() *Stmt {
	// vals := st.pattern
	// vals[methodKey] = mDelete
	// st.pattern = vals
	return st
}

// Select -
func (st *Stmt) Select() *Stmt {
	// vals := st.pattern
	// vals[methodKey] = mSelect
	// st.pattern = vals
	return st
}

// Columns -
func (st *Stmt) Columns(args interface{}) *Stmt {
	//vals := st.pattern

	switch args.(type) {
	case []string:
		//vals[columnsKey] = args.([]string)
	case string:
		//vals[columnsKey] = args.(string)
	case []struct{}:
		//arr := args.([]struct{})
		//strVals := []string{}
		//for _, arg := range arr {
		//strVals = append(strVals, parse(arg)...)
		//}
		//vals[columnsKey] = strVals
	case struct{}:
		//vals[columnsKey] = parseStruct(args)
	default:
		panic("unsupported argument type")
	}

	//st.pattern = vals
	return st
}

// From -
func (st *Stmt) From(args interface{}) *Stmt {
	//vals := st.pattern

	switch args.(type) {
	case []string:
		//vals[fromKey] = args.([]string)
	case string:
		//vals[fromKey] = []string{args.(string)}
	default:
		panic("unsupported argument type")
	}

	//st.pattern = vals
	return st
}

// Into -
func (st *Stmt) Into() *Stmt {
	//vals := st.pattern
	//vals[intoKey] = mInto
	return st
}

// func parseStruct(arg interface{}) map[string]interface{} {
// 	v := reflect.ValueOf(arg)
// 	t := reflect.TypeOf(arg)

// 	res := map[string]interface{}{}

// 	for i := 0; i < v.NumField(); i++ {
// 		val := v.Field(i)
// 		field := t.Field(i)
// 		switch val.Kind() {
// 		case reflect.Struct:
// 			nested := parseStruct(val.Interface())
// 			for k, v := range nested {
// 				res[k] = v
// 			}
// 		default:
// 			tag := field.Tag.Get(tagKey)
// 			res[tag] = val
// 		}
// 	}

// 	return res
// }
