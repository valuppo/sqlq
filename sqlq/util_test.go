package sqlq

import "testing"

func TestCheckSameColumns(t *testing.T) {
	tables := []struct{
		Input []string
		Output bool
	}{
		{
			[]string{},
			false,
		},
		{
			[]string{"column1", "column2", "column3"},
			false,
		},
		{
			[]string{"column1", "column1", "column3"},
			true,
		},
		{
			[]string{"column1", "column2", "column1"},
			true,
		},
		{
			[]string{"column2", "column2", "column1"},
			true,
		},
		{
			[]string{"column1", "column2", "column2"},
			true,
		},
		{
			[]string{"column3", "column2", "column3"},
			true,
		},
		{
			[]string{"column1", "column3", "column3"},
			true,
		},
		{
			[]string{"column1", "column1", "column1"},
			true,
		},
	}
	for _, table := range tables{
		result := checkSameColumns(table.Input)
		if result != table.Output{
			t.Errorf("Expected %v got %v", table.Output, result)
		}
	}
}