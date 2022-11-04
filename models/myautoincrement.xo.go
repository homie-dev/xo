// Package models contains the types for schema 'xodb'.
package models

// Code generated by xo. DO NOT EDIT.

// MyAutoIncrement represents a row from '[custom my_auto_increment]'.
type MyAutoIncrement struct {
	TableName string // table_name
}

// MyAutoIncrements runs a custom query, returning results as MyAutoIncrement.
func MyAutoIncrements(db XODB, schema string) ([]*MyAutoIncrement, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`table_name ` +
		`FROM information_schema.columns ` +
		`WHERE extra = 'auto_increment' AND table_schema = ?`

	// run query
	XOLog(sqlstr, schema)
	q, err := db.Query(sqlstr, schema)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*MyAutoIncrement{}
	for q.Next() {
		mai := MyAutoIncrement{}

		// scan
		err = q.Scan(&mai.TableName)
		if err != nil {
			return nil, err
		}

		res = append(res, &mai)
	}

	return res, nil
}
