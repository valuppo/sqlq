package sqlq

import (
	"reflect"
	"testing"
)

func TestUpdate(t *testing.T) {
	tables := []struct {
		Table  string
		Output *UpdateBuilder
	}{
		{
			"users",
			Update("users"),
		},
		{
			"",
			Update(""),
		},
	}
	for _, table := range tables {
		result := Update(table.Table)
		if !reflect.DeepEqual(result, table.Output) {
			t.Errorf("Expected %v got %v", table.Output, result)
		}
	}
}

func TestUpdateBuilder_Set(t *testing.T) {
	tables := []struct {
		Column string
		Value  string
		Output *UpdateBuilder
	}{
		{
			"name",
			"sqlq",
			Update("").Set("name", "sqlq"),
		},
		{
			"email",
			"sqlq@valuppo.com",
			Update("").Set("email", "sqlq@valuppo.com"),
		},
		{
			"",
			"",
			Update("").Set("", ""),
		},
	}
	for _, table := range tables {
		result := Update("").Set(table.Column, table.Value)
		if !reflect.DeepEqual(result, table.Output) {
			t.Errorf("Expected %v got %v", table.Output, result)
		}
	}
}

func TestUpdateBuilder_SetMultiple(t *testing.T) {
	tables := []struct {
		Columns []string
		Values  []string
		Output  *UpdateBuilder
	}{
		{
			[]string{"name", "email"},
			[]string{"sqlq", "sqlq@valuppo.com"},
			Update("").SetMultiple([]string{"name", "email"}, []string{"sqlq", "sqlq@valuppo.com"}),
		},
		{
			[]string{"name"},
			[]string{"sqlq"},
			Update("").SetMultiple([]string{"name"}, []string{"sqlq"}),
		},
		{
			[]string{},
			[]string{},
			Update("").SetMultiple([]string{}, []string{}),
		},
	}
	for _, table := range tables {
		result := Update("").SetMultiple(table.Columns, table.Values)
		if !reflect.DeepEqual(result, table.Output) {
			t.Errorf("Expected %v got %v", table.Output, result)
		}
	}
}

func TestUpdateBuilder_Where(t *testing.T) {
	tables := []struct {
		Column   string
		Operator string
		Value    string
		Output   *UpdateBuilder
	}{
		{
			"id",
			"=",
			"1",
			Update("").Where("id", "=", "1"),
		},
		{
			"name",
			"=",
			"sqlq",
			Update("").Where("name", "=", "sqlq"),
		},
		{
			"",
			"",
			"",
			Update("").Where("", "", ""),
		},
	}
	for _, table := range tables {
		result := Update("").Where(table.Column, table.Operator, table.Value)
		if !reflect.DeepEqual(result, table.Output) {
			t.Errorf("Expected %v got %v", table.Output, result)
		}
	}
}

func TestUpdateBuilder_WhereOr(t *testing.T) {
	tables := []struct {
		Column   string
		Operator string
		Value    string
		Output   *UpdateBuilder
	}{
		{
			"id",
			"=",
			"1",
			Update("").WhereOr("id", "=", "1"),
		},
		{
			"name",
			"=",
			"sqlq",
			Update("").WhereOr("name", "=", "sqlq"),
		},
		{
			"",
			"",
			"",
			Update("").WhereOr("", "", ""),
		},
	}
	for _, table := range tables {
		result := Update("").WhereOr(table.Column, table.Operator, table.Value)
		if !reflect.DeepEqual(result, table.Output) {
			t.Errorf("Expected %v got %v", table.Output, result)
		}
	}
}

func TestUpdateBuilder_Sql(t *testing.T) {
	tables := []struct {
		Table                 string
		Columns               []string
		Values                []string
		ConditionColumnsAnd   []string
		ConditionOperatorsAnd []string
		ConditionValuesAnd    []string
		ConditionColumnsOr    []string
		ConditionOperatorsOr  []string
		ConditionValuesOr     []string
		Output                string
		Error                 error
	}{
		{
			"users",
			[]string{"id", "name", "email"},
			[]string{"1", "sqlq", "sqlq@valuppo.com"},
			[]string{"id", "name"},
			[]string{">", "LIKE"},
			[]string{"1", "s%"},
			[]string{"email", "created_at"},
			[]string{"LIKE", ">="},
			[]string{"%@valuppo.com", "CURDATE()"},
			"UPDATE users SET id = '1', name = 'sqlq', email = 'sqlq@valuppo.com' WHERE id > 1 AND name LIKE 's%' " +
				"OR email LIKE '%@valuppo.com' OR created_at >= 'CURDATE()'",
			nil,
		},
		{
			"users",
			[]string{"id", "name", "email"},
			[]string{"1", "sqlq", "sqlq@valuppo.com"},
			[]string{"id", "name"},
			[]string{">", "LIKE"},
			[]string{"1", "s%"},
			[]string{},
			[]string{},
			[]string{},
			"UPDATE users SET id = '1', name = 'sqlq', email = 'sqlq@valuppo.com' WHERE id > 1 AND name LIKE 's%'",
			nil,
		},
		{
			"users",
			[]string{"id", "name", "email"},
			[]string{"1", "sqlq", "sqlq@valuppo.com"},
			[]string{},
			[]string{},
			[]string{},
			[]string{"email", "created_at"},
			[]string{"LIKE", ">="},
			[]string{"%@valuppo.com", "CURDATE()"},
			"UPDATE users SET id = '1', name = 'sqlq', email = 'sqlq@valuppo.com' WHERE email LIKE '%@valuppo.com' " +
				"OR created_at >= 'CURDATE()'",
			nil,
		},
		{
			"users",
			[]string{"id", "name", "email"},
			[]string{"1", "sqlq", "sqlq@valuppo.com"},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			"UPDATE users SET id = '1', name = 'sqlq', email = 'sqlq@valuppo.com'",
			nil,
		},
		{
			"",
			[]string{"id", "name", "email"},
			[]string{"1", "sqlq", "sqlq@valuppo.com"},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			"",
			errUpdateEmptyTables,
		},
		{
			"users",
			[]string{},
			[]string{"1", "sqlq", "sqlq@valuppo.com"},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			"",
			errUpdateEmptyColumns,
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
			"",
			errUpdateEmptyValues,
		},
		{
			"users",
			[]string{"id", "name", "email", ""},
			[]string{"1", "sqlq", "sqlq@valuppo.com", "test"},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			"",
			errUpdateColumnsValuesDiffLen,
		},
		{
			"users",
			[]string{"id", "name", "email", "test"},
			[]string{"1", "sqlq", "sqlq@valuppo.com", ""},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			"",
			errUpdateColumnsValuesDiffLen,
		},
		{
			"users",
			[]string{"id", "name", "email", "id"},
			[]string{"1", "sqlq", "sqlq@valuppo.com", "2"},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			"",
			errUpdateColumnsSame,
		},
	}
	for _, table := range tables {
		update := Update(table.Table)
		for i := 0; i < len(table.Columns) && i < len(table.Values); i++ {
			update.Set(table.Columns[i], table.Values[i])
		}
		for i := 0; i < len(table.ConditionColumnsAnd) && i < len(table.ConditionOperatorsAnd) && i < len(table.ConditionValuesAnd); i++ {
			update.Where(table.ConditionColumnsAnd[i], table.ConditionOperatorsAnd[i], table.ConditionValuesAnd[i])
		}
		for i := 0; i < len(table.ConditionColumnsOr) && i < len(table.ConditionOperatorsOr) && i < len(table.ConditionValuesOr); i++ {
			update.Where(table.ConditionColumnsOr[i], table.ConditionOperatorsOr[i], table.ConditionValuesOr[i])
		}
		result, err := update.Sql()
		if result != table.Output && err != table.Error {
			t.Errorf("Expected %v\nGot %v\n", table.Output, result)
			t.Errorf("Expected error %v\nGot %v\n", table.Error, err)
		}
	}
}

func TestUpdateBuilder_Sql2(t *testing.T) {
	tables := []struct {
		Table                 string
		Columns               []string
		Values                []string
		ConditionColumnsAnd   []string
		ConditionOperatorsAnd []string
		ConditionValuesAnd    []string
		ConditionColumnsOr    []string
		ConditionOperatorsOr  []string
		ConditionValuesOr     []string
		Output                string
		Error                 error
	}{
		{
			"users",
			[]string{"id", "name", "email"},
			[]string{"1", "sqlq", "sqlq@valuppo.com"},
			[]string{"id", "name"},
			[]string{">", "LIKE"},
			[]string{"1", "s%"},
			[]string{"email", "created_at"},
			[]string{"LIKE", ">="},
			[]string{"%@valuppo.com", "CURDATE()"},
			"UPDATE users SET id = '1', name = 'sqlq', email = 'sqlq@valuppo.com' WHERE id > 1 AND name LIKE 's%' " +
				"OR email LIKE '%@valuppo.com' OR created_at >= 'CURDATE()'",
			nil,
		},
		{
			"users",
			[]string{"id", "name", "email"},
			[]string{"1", "sqlq", "sqlq@valuppo.com"},
			[]string{"id", "name"},
			[]string{">", "LIKE"},
			[]string{"1", "s%"},
			[]string{},
			[]string{},
			[]string{},
			"UPDATE users SET id = '1', name = 'sqlq', email = 'sqlq@valuppo.com' WHERE id > 1 AND name LIKE 's%'",
			nil,
		},
		{
			"users",
			[]string{"id", "name", "email"},
			[]string{"1", "sqlq", "sqlq@valuppo.com"},
			[]string{},
			[]string{},
			[]string{},
			[]string{"email", "created_at"},
			[]string{"LIKE", ">="},
			[]string{"%@valuppo.com", "CURDATE()"},
			"UPDATE users SET id = '1', name = 'sqlq', email = 'sqlq@valuppo.com' WHERE email LIKE '%@valuppo.com' " +
				"OR created_at >= 'CURDATE()'",
			nil,
		},
		{
			"users",
			[]string{"id", "name", "email"},
			[]string{"1", "sqlq", "sqlq@valuppo.com"},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			"UPDATE users SET id = '1', name = 'sqlq', email = 'sqlq@valuppo.com'",
			nil,
		},
		{
			"",
			[]string{"id", "name", "email"},
			[]string{"1", "sqlq", "sqlq@valuppo.com"},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			"",
			errUpdateEmptyTables,
		},
		{
			"users",
			[]string{},
			[]string{"1", "sqlq", "sqlq@valuppo.com"},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			"",
			errUpdateEmptyColumns,
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
			"",
			errUpdateEmptyValues,
		},
		{
			"users",
			[]string{"id", "name", "email", ""},
			[]string{"1", "sqlq", "sqlq@valuppo.com", "test"},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			"",
			errUpdateColumnsValuesDiffLen,
		},
		{
			"users",
			[]string{"id", "name", "email", "test"},
			[]string{"1", "sqlq", "sqlq@valuppo.com", ""},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			"",
			errUpdateColumnsValuesDiffLen,
		},
		{
			"users",
			[]string{"id", "name", "email", "id"},
			[]string{"1", "sqlq", "sqlq@valuppo.com", "2"},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			"",
			errUpdateColumnsSame,
		},
	}
	for _, table := range tables {
		update := Update(table.Table)
		update.SetMultiple(table.Columns, table.Values)
		for i := 0; i < len(table.ConditionColumnsAnd) && i < len(table.ConditionOperatorsAnd) && i < len(table.ConditionValuesAnd); i++ {
			update.Where(table.ConditionColumnsAnd[i], table.ConditionOperatorsAnd[i], table.ConditionValuesAnd[i])
		}
		for i := 0; i < len(table.ConditionColumnsOr) && i < len(table.ConditionOperatorsOr) && i < len(table.ConditionValuesOr); i++ {
			update.Where(table.ConditionColumnsOr[i], table.ConditionOperatorsOr[i], table.ConditionValuesOr[i])
		}
		result, err := update.Sql()
		if result != table.Output && err != table.Error {
			t.Errorf("Expected %v\nGot %v\n", table.Output, result)
			t.Errorf("Expected error %v\nGot %v\n", table.Error, err)
		}
	}
}
