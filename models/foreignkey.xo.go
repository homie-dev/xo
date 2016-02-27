// Package models contains the types for schema 'public'.
package models

// GENERATED BY XO. DO NOT EDIT.

// ForeignKey represents a foreign key.
type ForeignKey struct {
	ForeignKeyName string // foreign_key_name
	ColumnName     string // column_name
	RefIndexName   string // ref_index_name
	RefTableName   string // ref_table_name
	RefColumnName  string // ref_column_name
	KeyID          int    // key_id
	SeqNo          int    // seq_no
	OnUpdate       string // on_update
	OnDelete       string // on_delete
	Match          string // match
}

// PgTableForeignKeys runs a custom query, returning results as ForeignKey.
func PgTableForeignKeys(db XODB, schema string, table string) ([]*ForeignKey, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`r.conname, ` + // ::varchar AS foreign_key_name
		`b.attname, ` + // ::varchar AS column_name
		`i.relname, ` + // ::varchar AS ref_index_name
		`c.relname, ` + // ::varchar AS ref_table_name
		`d.attname, ` + // ::varchar AS ref_column_name
		`0, ` + // ::integer AS key_id
		`0, ` + // ::integer AS seq_no
		`'', ` + // ::varchar AS on_update
		`'', ` + // ::varchar AS on_delete
		`'' ` + // ::varchar AS match
		`FROM pg_constraint r ` +
		`JOIN ONLY pg_class a ON a.oid = r.conrelid ` +
		`JOIN ONLY pg_attribute b ON b.attisdropped = false AND b.attnum = ANY(r.conkey) AND b.attrelid = r.conrelid ` +
		`JOIN ONLY pg_class i on i.oid = r.conindid ` +
		`JOIN ONLY pg_class c on c.oid = r.confrelid ` +
		`JOIN ONLY pg_attribute d ON d.attisdropped = false AND d.attnum = ANY(r.confkey) AND d.attrelid = r.confrelid ` +
		`JOIN ONLY pg_namespace n ON n.oid = r.connamespace ` +
		`WHERE r.contype = 'f' AND n.nspname = $1 AND a.relname = $2 ` +
		`ORDER BY r.conname, b.attname`

	// run query
	XOLog(sqlstr, schema, table)
	q, err := db.Query(sqlstr, schema, table)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*ForeignKey{}
	for q.Next() {
		fk := ForeignKey{}

		// scan
		err = q.Scan(&fk.ForeignKeyName, &fk.ColumnName, &fk.RefIndexName, &fk.RefTableName, &fk.RefColumnName, &fk.KeyID, &fk.SeqNo, &fk.OnUpdate, &fk.OnDelete, &fk.Match)
		if err != nil {
			return nil, err
		}

		res = append(res, &fk)
	}

	return res, nil
}

// MyTableForeignKeys runs a custom query, returning results as ForeignKey.
func MyTableForeignKeys(db XODB, schema string, table string) ([]*ForeignKey, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`constraint_name AS foreign_key_name, ` +
		`column_name AS column_name, ` +
		`referenced_table_name AS ref_table_name, ` +
		`referenced_column_name AS ref_column_name ` +
		`FROM information_schema.key_column_usage ` +
		`WHERE referenced_table_name IS NOT NULL AND table_schema = ? AND table_name = ?`

	// run query
	XOLog(sqlstr, schema, table)
	q, err := db.Query(sqlstr, schema, table)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*ForeignKey{}
	for q.Next() {
		fk := ForeignKey{}

		// scan
		err = q.Scan(&fk.ForeignKeyName, &fk.ColumnName, &fk.RefTableName, &fk.RefColumnName)
		if err != nil {
			return nil, err
		}

		res = append(res, &fk)
	}

	return res, nil
}

// SqTableForeignKeys runs a custom query, returning results as ForeignKey.
func SqTableForeignKeys(db XODB, table string) ([]*ForeignKey, error) {
	var err error

	// sql query
	var sqlstr = `PRAGMA foreign_key_list(` + table + `)`

	// run query
	XOLog(sqlstr)
	q, err := db.Query(sqlstr, table)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*ForeignKey{}
	for q.Next() {
		fk := ForeignKey{}

		// scan
		err = q.Scan(&fk.KeyID, &fk.SeqNo, &fk.RefTableName, &fk.ColumnName, &fk.RefColumnName, &fk.OnUpdate, &fk.OnDelete, &fk.Match)
		if err != nil {
			return nil, err
		}

		res = append(res, &fk)
	}

	return res, nil
}

// OrTableForeignKeys runs a custom query, returning results as ForeignKey.
func OrTableForeignKeys(db XODB, schema string, table string) ([]*ForeignKey, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`a.constraint_name AS foreign_key_name, ` +
		`a.column_name, ` +
		`r.constraint_name AS ref_index_name, ` +
		`r.table_name AS ref_table_name ` +
		`FROM all_cons_columns a ` +
		`JOIN all_constraints c ON a.owner = c.owner AND a.constraint_name = c.constraint_name ` +
		`JOIN all_constraints r ON c.r_owner = r.owner AND c.r_constraint_name = r.constraint_name ` +
		`WHERE c.constraint_type = 'R' AND a.owner = UPPER(:1) AND a.table_name = UPPER(:2)`

	// run query
	XOLog(sqlstr, schema, table)
	q, err := db.Query(sqlstr, schema, table)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*ForeignKey{}
	for q.Next() {
		fk := ForeignKey{}

		// scan
		err = q.Scan(&fk.ForeignKeyName, &fk.ColumnName, &fk.RefIndexName, &fk.RefTableName)
		if err != nil {
			return nil, err
		}

		res = append(res, &fk)
	}

	return res, nil
}
