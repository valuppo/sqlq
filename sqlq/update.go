package sqlq

import (
	"errors"
	"strings"
)

var errUpdateEmptyTables = errors.New("\nTable Name is required to do Update Operation\nUse Update() function to specify Table Name")
var errUpdateEmptyColumns = errors.New("\nColumns is required to do Update Operation\nUse Set() function to specify Columns")
var errUpdateEmptyValues = errors.New("\nValues is required to do Update Operation\nUse Set() function to specify Values")

var errUpdateColumnsValuesDiffLen = errors.New("\nLength of Columns and Values must be the same")
var errUpdateColumnsSame = errors.New("You have same Columns in your query")

type UpdateBuilder struct {
	table         string
	columns       []string
	values        []string
	conditionsAnd []string
	conditionsOr  []string
}

func (ub *UpdateBuilder) Set(column string, value string) *UpdateBuilder {
	if column != "" {
		ub.columns = append(ub.columns, column)
	}
	if value != "" {
		ub.values = append(ub.values, "'"+value+"'")
	}
	return ub
}

func (ub *UpdateBuilder) SetMultiple(columns []string, values []string) *UpdateBuilder {
	for _, column := range columns {
		if column != "" {
			ub.columns = append(ub.columns, column)
		}
	}
	for _, value := range values {
		if value != "" {
			ub.values = append(ub.values, "'"+value+"'")
		}
	}
	return ub
}

func (ub *UpdateBuilder) Where(column string, operator string, value string) *UpdateBuilder {
	if column != "" && operator != "" && value != "" {
		ub.conditionsAnd = append(ub.conditionsAnd, column+" "+operator+" '"+value+"'")
	}
	return ub
}

func (ub *UpdateBuilder) WhereOr(column string, operator string, value string) *UpdateBuilder {
	if column != "" && operator != "" && value != "" {
		ub.conditionsOr = append(ub.conditionsOr, column+" "+operator+" '"+value+"'")
	}
	return ub
}

func (ub *UpdateBuilder) Sql() (string, error) {
	var err error

	if ub.table == "" {
		err = errUpdateEmptyTables
		return "", err
	} else if len(ub.columns) <= 0 {
		err = errUpdateEmptyColumns
		return "", err
	} else if len(ub.values) <= 0 {
		err = errUpdateEmptyValues
		return "", err
	} else if len(ub.columns) != len(ub.values) {
		err = errUpdateColumnsValuesDiffLen
		return "", err
	} else if checkSameColumns(ub.columns) {
		err = errUpdateColumnsSame
		return "", err
	}

	sqlSet := make([]string, len(ub.columns))
	for i := 0; i < len(ub.columns) && i < len(ub.values); i++ {
		sqlSet[i] = ub.columns[i] + " = " + ub.values[i]
	}

	sql := "UPDATE " + ub.table + " SET " + strings.Join(sqlSet, ", ")
	if len(ub.conditionsAnd) > 0 || len(ub.conditionsOr) > 0 {
		sql += " WHERE "
	}
	if len(ub.conditionsAnd) > 0 {
		sql += strings.Join(ub.conditionsAnd, " AND ")
	}
	if len(ub.conditionsOr) > 0 {
		if len(ub.conditionsAnd) > 0 {
			sql += " OR"
		}
		sql += " " + strings.Join(ub.conditionsOr, " OR ")
	}
	return sql, err
}

func Update(table string) *UpdateBuilder {
	return &UpdateBuilder{
		table: table,
	}
}
