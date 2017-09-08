package sqlq

import (
	"errors"
	"strconv"
	"strings"
)

var errSelectEmptyTable = errors.New("\nTable Name is required to do Select Operation\nUse From() function to specify Table Name")
var errSelectEmptyColumns = errors.New("\nColumns is required to do Select Operation\nUse Select() or Columns() function to specify Columns")

var errSelectLimitNegative = errors.New("\nLimit can't be negative value")
var errSelectColumnsSame = errors.New("You have same Columns in your query")

type SelectBuilder struct {
	table         string
	columns       []string
	conditionsAnd []string
	conditionsOr  []string
	order         []string
	limit         int
}

func (sb *SelectBuilder) From(table string) *SelectBuilder {
	sb.table = table
	return sb
}

func (sb *SelectBuilder) Columns(columns ...string) *SelectBuilder {
	sb.columns = append(sb.columns, columns...)
	return sb
}

func (sb *SelectBuilder) Where(column string, operator string, value string) *SelectBuilder {
	if column != "" && operator != "" && value != "" {
		sb.conditionsAnd = append(sb.conditionsAnd, column+" "+operator+" '"+value+"'")
	}
	return sb
}

func (sb *SelectBuilder) WhereOr(column string, operator string, value string) *SelectBuilder {
	if column != "" && operator != "" && value != "" {
		sb.conditionsOr = append(sb.conditionsOr, column+" "+operator+" '"+value+"'")
	}
	return sb
}

func (sb *SelectBuilder) OrderBy(column string, order string) *SelectBuilder {
	sb.order = append(sb.order, column+" "+order)
	return sb
}

func (sb *SelectBuilder) Limit(limit int) *SelectBuilder {
	sb.limit = limit
	return sb
}

func (sb *SelectBuilder) Sql() (string, error) {
	var err error

	if sb.table == "" {
		err = errSelectEmptyTable
		return "", err
	} else if len(sb.columns) <= 0 {
		err = errSelectEmptyColumns
		return "", err
	} else if sb.limit < 0 {
		err = errSelectLimitNegative
		return "", err
	} else if checkSameColumns(sb.columns) {
		err = errSelectColumnsSame
		return "", err
	}

	sql := "SELECT " + strings.Join(sb.columns, ", ") + " FROM " + sb.table
	if len(sb.conditionsAnd) > 0 || len(sb.conditionsOr) > 0 {
		sql += " WHERE "
	}
	if len(sb.conditionsAnd) > 0 {
		sql += strings.Join(sb.conditionsAnd, " AND ")
	}
	if len(sb.conditionsOr) > 0 {
		if len(sb.conditionsAnd) > 0 {
			sql += " OR"
		}
		sql += " " + strings.Join(sb.conditionsOr, " OR ")
	}
	if len(sb.order) > 0 {
		sql += " ORDER BY " + strings.Join(sb.order, ", ")
	}
	if sb.limit > 0 {
		sql += " LIMIT " + strconv.Itoa(sb.limit)
	}
	return sql, err
}

func Select(columns ...string) *SelectBuilder {
	sb := &SelectBuilder{}
	sb.columns = append(sb.columns, columns...)
	return sb
}
