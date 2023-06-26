package query

func Delete(table string) *Sql {
	return &Sql{
		query: `DELETE FROM ` + table,
		table: table,
	}
}
