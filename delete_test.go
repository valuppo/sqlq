package sqlq

import (
	"reflect"
	"testing"
)

func TestDelete(t *testing.T) {
	tables := []struct {
		Output *DeleteBuilder
	}{
		{
			Delete(),
		},
	}
	for _, table := range tables {
		result := Delete()
		if !reflect.DeepEqual(result, table.Output) {
			t.Errorf("Expected %v got %v", table.Output, result)
		}
	}
}

func TestDeleteBuilder_From(t *testing.T) {
	tables := []struct {
		Input  string
		Output *DeleteBuilder
	}{
		{
			"users",
			Delete().From("users"),
		},
		{
			"test",
			Delete().From("test"),
		},
		{
			"",
			Delete().From(""),
		},
	}
	for _, table := range tables {
		result := Delete().From(table.Input)
		if !reflect.DeepEqual(result, table.Output) {
			t.Errorf("Expected %v got %v", table.Output, result)
		}
	}
}

func TestDeleteBuilder_Where(t *testing.T) {
	tables := []struct {
		Column   string
		Operator string
		Value    string
		Output   *DeleteBuilder
	}{
		{
			"id",
			"=",
			"1",
			Delete().Where("id", "=", "1"),
		},
		{
			"name",
			"=",
			"sqlq",
			Delete().Where("name", "=", "sqlq"),
		},
		{
			"",
			"",
			"",
			Delete().Where("", "", ""),
		},
	}
	for _, table := range tables {
		result := Delete().Where(table.Column, table.Operator, table.Value)
		if !reflect.DeepEqual(result, table.Output) {
			t.Errorf("Expected %v got %v", table.Output, result)
		}
	}
}

func TestDeleteBuilder_WhereOr(t *testing.T) {
	tables := []struct {
		Column   string
		Operator string
		Value    string
		Output   *DeleteBuilder
	}{
		{
			"id",
			"=",
			"1",
			Delete().WhereOr("id", "=", "1"),
		},
		{
			"name",
			"=",
			"sqlq",
			Delete().WhereOr("name", "=", "sqlq"),
		},
		{
			"",
			"",
			"",
			Delete().WhereOr("", "", ""),
		},
	}
	for _, table := range tables {
		result := Delete().WhereOr(table.Column, table.Operator, table.Value)
		if !reflect.DeepEqual(result, table.Output) {
			t.Errorf("Expected %v got %v", table.Output, result)
		}
	}
}

func TestDeleteBuilder_Sql(t *testing.T) {
	tables := []struct {
		Table                 string
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
			[]string{"id", "name"},
			[]string{"=", "LIKE"},
			[]string{"1", "%sqlq%"},
			[]string{"email", "created_at"},
			[]string{"LIKE", ">"},
			[]string{"sqlq@%", "CURDATE()"},
			"DELETE FROM users WHERE id = 1 AND name LIKE '%sqlq' OR email LIKE 'sqlq@%' OR created_at > CURDATE()",
			nil,
		},
		{
			"users",
			[]string{},
			[]string{},
			[]string{},
			[]string{"email", "created_at"},
			[]string{"LIKE", ">"},
			[]string{"sqlq@%", "CURDATE()"},
			"DELETE FROM users WHERE email LIKE 'sqlq@%' OR created_at > CURDATE()",
			nil,
		},
		{
			"users",
			[]string{"id", "name"},
			[]string{"=", "LIKE"},
			[]string{"1", "%sqlq%"},
			[]string{},
			[]string{},
			[]string{},
			"DELETE FROM users WHERE id = 1 AND name LIKE '%sqlq'",
			nil,
		},
		{
			"users",
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			"DELETE FROM users",
			nil,
		},
		{
			"",
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			[]string{},
			"",
			errDeleteEmptyTable,
		},
	}
	for _, table := range tables {
		delete := Delete().From(table.Table)
		for i := 0; i < len(table.ConditionColumnsAnd) && i < len(table.ConditionOperatorsAnd) && i < len(table.ConditionValuesAnd); i++{
			delete.Where(table.ConditionColumnsAnd[i], table.ConditionOperatorsAnd[i], table.ConditionValuesAnd[i])
		}
		for i := 0; i < len(table.ConditionColumnsOr) && i < len(table.ConditionOperatorsOr) && i < len(table.ConditionValuesOr); i++{
			delete.Where(table.ConditionColumnsOr[i], table.ConditionOperatorsOr[i], table.ConditionValuesOr[i])
		}
		result, err := delete.Sql()
		if result != table.Output && err != table.Error {
			t.Errorf("Expected %v\nGot %v\n", table.Output, result)
			t.Errorf("Expected error %v\nGot %v\n", table.Error, err)
		}
	}
}
