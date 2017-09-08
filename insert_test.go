package sqlq

import (
	"testing"
	"reflect"
)

func TestInsert(t *testing.T) {
	tables := []struct{
		Output *InsertBuilder
	}{
		{
			Insert(),
		},
	}
	for _, table := range tables{
		result := Insert()
		if !reflect.DeepEqual(result, table.Output){
			t.Errorf("Expected %v got %v", table.Output, result)
		}
	}
}

func TestInsertBuilder_Into(t *testing.T) {
	tables := []struct{
		Input string
		Output *InsertBuilder
	}{
		{
			"users",
			Insert().Into("users"),
		},
		{
			"test",
			Insert().Into("test"),
		},
		{
			"",
			Insert().Into(""),
		},
	}
	for _, table := range tables{
		result := Insert().Into(table.Input)
		if !reflect.DeepEqual(result, table.Output){
			t.Errorf("Expected %v got %v", table.Output, result)
		}
	}
}

func TestInsertBuilder_Columns(t *testing.T) {
	tables := []struct{
		Input []string
		Output *InsertBuilder
	}{
		{
			[]string{"email", "name", "password"},
			Insert().Columns([]string{"email", "name", "password"}...),
		},
		{
			[]string{"email"},
			Insert().Columns([]string{"email"}...),
		},
		{
			[]string{},
			Insert().Columns([]string{}...),
		},
	}
	for _, table := range tables{
		result := Insert().Columns(table.Input...)
		if !reflect.DeepEqual(result, table.Output){
			t.Errorf("Expected %v got %v", table.Output, result)
		}
	}
}

func TestInsertBuilder_Values(t *testing.T) {
	tables := []struct{
		Input []string
		Output *InsertBuilder
	}{
		{
			[]string{"sqlq@valuppo.com", "sqlq", "sqlq_pass"},
			Insert().Columns([]string{"sqlq@valuppo.com", "sqlq", "sqlq_pass"}...),
		},
		{
			[]string{"sqlq@valuppo.com"},
			Insert().Columns([]string{"sqlq@valuppo.com"}...),
		},
		{
			[]string{},
			Insert().Columns([]string{}...),
		},
	}
	for _, table := range tables{
		result := Insert().Columns(table.Input...)
		if !reflect.DeepEqual(result, table.Output){
			t.Errorf("Expected %v got %v", table.Output, result)
		}
	}
}

func TestInsertBuilder_Sql(t *testing.T) {
	tables := []struct{
		Table string
		Columns []string
		Values []string
		Output string
		Error error
	}{
		{
			"users",
			[]string{"email", "name", "password"},
			[]string{"sqlq@valuppo.com", "sqlq", "sqlq_pass"},
			"INSERT INTO users (email, name, password) VALUES ('sqlq@valuppo.com', 'sqlq', 'sqlq_pass')",
			nil,
		},
		{
			"",
			[]string{"email", "name", "password"},
			[]string{"sqlq@valuppo.com", "sqlq", "sqlq_pass"},
			"",
			errInsertEmptyTable,
		},
		{
			"users",
			[]string{""},
			[]string{"sqlq@valuppo.com", "sqlq", "sqlq_pass"},
			"",
			errInsertEmptyColumns,
		},
		{
			"users",
			[]string{"email", "name", "password"},
			[]string{""},
			"",
			errInsertEmptyValues,
		},
		{
			"users",
			[]string{"email", "name", "password"},
			[]string{"'sqlq@valuppo.com'", "'sqlq'", "'sqlq_pass'", "'test'"},
			"",
			errInsertColumnsValuesDiffLen,
		},
		{
			"users",
			[]string{"email", "name", "password", "test"},
			[]string{"sqlq@valuppo.com", "sqlq", "sqlq_pass"},
			"",
			errInsertColumnsValuesDiffLen,
		},
		{
			"users",
			[]string{"email", "name", "password", "email"},
			[]string{"sqlq@valuppo.com", "sqlq", "sqlq_pass", "test"},
			"",
			errInsertColumnsSame,
		},
	}
	for _, table := range tables{
		result, err := Insert().Into(table.Table).Columns(table.Columns...).Values(table.Values...).Sql()
		if !reflect.DeepEqual(result, table.Output){
			t.Errorf("Expected %v\nGot %v\n", table.Output, result)
			t.Errorf("Expected error %v\nGot %v\n", table.Error, err)
		}
	}
}