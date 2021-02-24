package query

func Delete(table string) *sql {
	return &sql{
		query: `DELETE FROM ` + table,
	}
}
