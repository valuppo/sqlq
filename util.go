package sqlq

func checkSameColumns(columns []string) bool {
	for i := 0; i < len(columns); i++ {
		for j := i + 1; j < len(columns); j++ {
			if columns[i] == columns[j] {
				return true
			}
		}
	}
	return false
}
