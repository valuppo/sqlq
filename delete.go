package sqlq

import (
	"errors"
	"strings"
)

var errDeleteEmptyTable = errors.New("\nTable Name is required to do Delete Operation\nUse From() function to specify Table Name")

type DeleteBuilder struct {
	table         string
	conditionsAnd []string
	conditionsOr  []string
}

func (db *DeleteBuilder) From(table string) *DeleteBuilder {
	db.table = table
	return db
}

func (db *DeleteBuilder) Where(column string, operator string, value string) *DeleteBuilder {
	if column != "" && operator != "" && value != "" {
		db.conditionsAnd = append(db.conditionsAnd, column+" "+operator+" '"+value+"'")
	}
	return db
}

func (db *DeleteBuilder) WhereOr(column string, operator string, value string) *DeleteBuilder {
	if column != "" && operator != "" && value != "" {
		db.conditionsOr = append(db.conditionsOr, column+" "+operator+" '"+value+"'")
	}
	return db
}

func (db *DeleteBuilder) Sql() (string, error) {
	var err error
	if db.table == "" {
		err = errDeleteEmptyTable
		return "", err
	}
	sql := "DELETE FROM " + db.table
	if len(db.conditionsAnd) > 0 || len(db.conditionsOr) > 0 {
		sql += " WHERE "
	}
	if len(db.conditionsAnd) > 0 {
		sql += strings.Join(db.conditionsAnd, " AND ")
	}
	if len(db.conditionsOr) > 0 {
		if len(db.conditionsAnd) > 0 {
			sql += " OR"
		}
		sql += " " + strings.Join(db.conditionsOr, " OR ")
	}
	return sql, err
}

func Delete() *DeleteBuilder {
	return &DeleteBuilder{}
}
