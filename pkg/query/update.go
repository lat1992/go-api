package query

func Update(table string) *Sql {
	return &Sql{
		query: `UPDATE ` + table,
		table: table,
	}
}

func (s *Sql) Set(columnsValues []ColumnValue) *Sql {
	s.query += ` SET `
	for i, cv := range columnsValues {
		s.query += cv.Column + ` = ` + s.getNextParams()
		s.args = append(s.args, cv.Value)
		if i < len(columnsValues)-1 {
			s.query += `,`
		}
		i++
	}
	return s
}
