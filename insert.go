package sqlq

import (
	"errors"
	"fmt"
	"strings"
)

var errInsertEmptyTable = errors.New("\nTable Name is required to do Insert Operation\nUse Into() function to specify Table Name")
var errInsertEmptyColumns = errors.New("\nColumns is required to do Insert Operation\nUse Columns() function to specify Columns")
var errInsertEmptyValues = errors.New("\nValues is required to do Insert Operation\nUse Values() function to specify Values")

var errInsertColumnsValuesDiffLen = errors.New("\nLength of Columns and Values must be the same")
var errInsertColumnsSame = errors.New("You have same Columns in your query")

type InsertBuilder struct {
	table   string
	columns []string
	values  []string
}

func (ib *InsertBuilder) Into(table string) *InsertBuilder {
	ib.table = table
	return ib
}

func (ib *InsertBuilder) Columns(columns ...string) *InsertBuilder {
	ib.columns = append(ib.columns, columns...)
	return ib
}

func (ib *InsertBuilder) Values(values ...string) *InsertBuilder {
	for _, val := range values {
		ib.values = append(ib.values, "'"+val+"'")
	}
	return ib
}

func (ib *InsertBuilder) Sql() (string, error) {
	var err error
	if ib.table == "" {
		err = errInsertEmptyTable
		return "", err
	} else if len(ib.columns) <= 0 {
		err = errInsertEmptyColumns
		return "", err
	} else if len(ib.values) <= 0 {
		err = errInsertEmptyValues
		return "", err
	} else if len(ib.columns) != len(ib.values) {
		err = errInsertColumnsValuesDiffLen
		return "", err
	} else if checkSameColumns(ib.columns) {
		err = errInsertColumnsSame
		return "", err
	}
	sql := "INSERT INTO " + ib.table + " (" + strings.Join(ib.columns, ", ") + ") VALUES (" + strings.Join(ib.values, ", ") + ")"
	return sql, err
}

func Insert() *InsertBuilder {
	return &InsertBuilder{}
}
