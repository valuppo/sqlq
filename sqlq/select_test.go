package sqlq

import (
	"reflect"
	"testing"
)

func TestSelect(t *testing.T) {
	tables := []struct {
		Columns []string
		Output  *SelectBuilder
	}{
		{
			[]string{},
			Select([]string{}...),
		},
		{
			[]string{"id", "name", "email"},
			Select([]string{"id", "name", "email"}...),
		},
	}
	for _, table := range tables {
		result := Select(table.Columns...)
		if !reflect.DeepEqual(result, table.Output) {
			t.Errorf("Expected %v got %v", table.Output, result)
		}
	}
}

func TestSelectBuilder_From(t *testing.T) {
	tables := []struct {
		Input  string
		Output *SelectBuilder
	}{
		{
			"",
			Select().From(""),
		},
		{
			"users",
			Select().From("users"),
		},
	}
	for _, table := range tables {
		result := Select().From(table.Input)
		if !reflect.DeepEqual(result, table.Output) {
			t.Errorf("Expected %v got %v", table.Output, result)
		}
	}
}

func TestSelectBuilder_Columns(t *testing.T) {
	tables := []struct {
		Columns []string
		Output  *SelectBuilder
	}{
		{
			[]string{},
			Select().Columns([]string{}...),
		},
		{
			[]string{"id", "name", "email"},
			Select().Columns([]string{"id", "name", "email"}...),
		},
	}
	for _, table := range tables {
		result := Select().Columns(table.Columns...)
		if !reflect.DeepEqual(result, table.Output) {
			t.Errorf("Expected %v got %v", table.Output, result)
		}
	}
}

func TestSelectBuilder_Where(t *testing.T) {
	tables := []struct {
		Column   string
		Operator string
		Value    string
		Output   *SelectBuilder
	}{
		{
			"id",
			"=",
			"1",
			Select("").Where("id", "=", "1"),
		},
		{
			"name",
			"=",
			"sqlq",
			Select("").Where("name", "=", "sqlq"),
		},
		{
			"",
			"",
			"",
			Select("").Where("", "", ""),
		},
	}
	for _, table := range tables {
		result := Select("").Where(table.Column, table.Operator, table.Value)
		if !reflect.DeepEqual(result, table.Output) {
			t.Errorf("Expected %v got %v", table.Output, result)
		}
	}
}

func TestSelectBuilder_WhereOr(t *testing.T) {
	tables := []struct {
		Column   string
		Operator string
		Value    string
		Output   *SelectBuilder
	}{
		{
			"id",
			"=",
			"1",
			Select("").WhereOr("id", "=", "1"),
		},
		{
			"name",
			"=",
			"sqlq",
			Select("").WhereOr("name", "=", "sqlq"),
		},
		{
			"",
			"",
			"",
			Select("").WhereOr("", "", ""),
		},
	}
	for _, table := range tables {
		result := Select("").WhereOr(table.Column, table.Operator, table.Value)
		if !reflect.DeepEqual(result, table.Output) {
			t.Errorf("Expected %v got %v", table.Output, result)
		}
	}
}

func TestSelectBuilder_Limit(t *testing.T) {
	tables := []struct {
		Limit  int
		Output *SelectBuilder
	}{
		{
			-1,
			Select("").Limit(-1),
		},
		{
			0,
			Select("").Limit(0),
		},
		{
			1,
			Select("").Limit(1),
		},
	}
	for _, table := range tables {
		result := Select("").Limit(table.Limit)
		if !reflect.DeepEqual(result, table.Output) {
			t.Errorf("Expected %v got %v", table.Output, result)
		}
	}
}

func TestSelectBuilder_OrderBy(t *testing.T) {
	tables := []struct {
		Column string
		Order  string
		Output *SelectBuilder
	}{
		{
			"user",
			"ASC",
			Select("").OrderBy("user", "ASC"),
		},
		{
			"email",
			"DESC",
			Select("").OrderBy("email", "DESC"),
		},
		{
			"created_at",
			"DESC",
			Select("").OrderBy("created_at", "DESC"),
		},
	}
	for _, table := range tables {
		result := Select("").OrderBy(table.Column, table.Order)
		if !reflect.DeepEqual(result, table.Output) {
			t.Errorf("Expected %v got %v", table.Output, result)
		}
	}
}

func TestSelectBuilder_Sql(t *testing.T) {
	tables := []struct {
		Table                 string
		Columns               []string
		ConditionColumnsAnd   []string
		ConditionOperatorsAnd []string
		ConditionValuesAnd    []string
		ConditionColumnsOr    []string
		ConditionOperatorsOr  []string
		ConditionValuesOr     []string
		OrderColumns          []string
		OrderOrders           []string
		Limit                 int
		Output                string
		Error                 error
	}{
		{
			"users",
			[]string{"id", "name", "email"},
			[]string{"id", "name"},
			[]string{">", "LIKE"},
			[]string{"1", "s%"},
			[]string{"email", "created_at"},
			[]string{"LIKE", ">="},
			[]string{"%@valuppo.com", "CURDATE()"},
			[]string{"created_at", "name"},
			[]string{"DESC", "ASC"},
			5,
			"SELECT id, name, email FROM users WHERE id > 1 AND name LIKE 's%' " +
				"OR email LIKE '%@valuppo.com' OR created_at >= 'CURDATE()' ORDER BY created_at DESC, name ASC LIMIT 5",
			nil,
		},
		{
			"users",
			[]string{"id", "name", "email"},
			[]string{"id", "name"},
			[]string{">", "LIKE"},
			[]string{"1", "s%"},
			[]string{"email", "created_at"},
			[]string{"LIKE", ">="},
			[]string{"%@valuppo.com", "CURDATE()"},
			[]string{"created_at", "name"},
			[]string{"DESC", "ASC"},
			0,
			"SELECT id, name, email FROM users WHERE id > 1 AND name LIKE 's%' " +
				"OR email LIKE '%@valuppo.com' OR created_at >= 'CURDATE()' ORDER BY created_at DESC, name ASC",
			nil,
		},
		{
			"users",
			[]string{"id", "name", "email"},
			[]string{"id", "name"},
			[]string{">", "LIKE"},
			[]string{"1", "s%"},
			[]string{"email", "created_at"},
			[]string{"LIKE", ">="},
			[]string{"%@valuppo.com", "CURDATE()"},
			[]string{"created_at"},
			[]string{"DESC"},
			0,
			"SELECT id, name, email FROM users WHERE id > 1 AND name LIKE 's%' " +
				"OR email LIKE '%@valuppo.com' OR created_at >= 'CURDATE()' ORDER BY created_at DESC",
			nil,
		},
		{
			"users",
			[]string{"id", "name", "email"},
			[]string{"id", "name"},
			[]string{">", "LIKE"},
			[]string{"1", "s%"},
			[]string{"email", "created_at"},
			[]string{"LIKE", ">="},
			[]string{"%@valuppo.com", "CURDATE()"},
			[]string{},
			[]string{},
			0,
			"SELECT id, name, email FROM users WHERE id > 1 AND name LIKE 's%' " +
				"OR email LIKE '%@valuppo.com' OR created_at >= 'CURDATE()'",
			nil,
		},
		{
			"users",
			[]string{"id", "name", "email"},
			[]string{"id", "name"},
			[]string{">", "LIKE"},
			[]string{"1", "s%"},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			0,
			"SELECT id, name, email FROM users WHERE id > 1 AND name LIKE 's%'",
			nil,
		},
		{
			"users",
			[]string{"id", "name", "email"},
			[]string{},
			[]string{},
			[]string{},
			[]string{"email", "created_at"},
			[]string{"LIKE", ">="},
			[]string{"%@valuppo.com", "CURDATE()"},
			[]string{},
			[]string{},
			0,
			"SELECT id, name, email FROM users WHERE email LIKE '%@valuppo.com' " +
				"OR created_at >= 'CURDATE()'",
			nil,
		},
		{
			"users",
			[]string{"id", "name", "email"},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			0,
			"SELECT id, name, email FROM users",
			nil,
		},
		{
			"",
			[]string{"id", "name", "email"},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			0,
			"",
			errSelectEmptyTable,
		},
		{
			"users",
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			0,
			"",
			errSelectEmptyColumns,
		},
		{
			"users",
			[]string{"id", "name", "email"},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			-1,
			"",
			errSelectLimitNegative,
		},
		{
			"users",
			[]string{"id", "name", "email", "id"},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			0,
			"",
			errSelectColumnsSame,
		},
	}
	for _, table := range tables {
		slct := Select().From(table.Table)
		for i := 0; i < len(table.Columns); i++ {
			slct.Columns(table.Columns[i])
		}
		for i := 0; i < len(table.ConditionColumnsAnd) && i < len(table.ConditionOperatorsAnd) && i < len(table.ConditionValuesAnd); i++ {
			slct.Where(table.ConditionColumnsAnd[i], table.ConditionOperatorsAnd[i], table.ConditionValuesAnd[i])
		}
		for i := 0; i < len(table.ConditionColumnsOr) && i < len(table.ConditionOperatorsOr) && i < len(table.ConditionValuesOr); i++ {
			slct.Where(table.ConditionColumnsOr[i], table.ConditionOperatorsOr[i], table.ConditionValuesOr[i])
		}
		for i := 0; i < len(table.OrderColumns) && i < len(table.OrderOrders); i++ {
			slct.OrderBy(table.OrderColumns[i], table.OrderOrders[i])
		}
		slct.Limit(table.Limit)
		result, err := slct.Sql()
		if result != table.Output && err != table.Error {
			t.Errorf("Expected %v\nGot %v\n", table.Output, result)
			t.Errorf("Expected error %v\nGot %v\n", table.Error, err)
		}
	}
}
