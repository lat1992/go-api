package query

import (
	"strings"
)

func Insert(table string) *Sql {
	return &Sql{
		query: `INSERT INTO ` + table,
		table: table,
	}
}

func (s *Sql) Columns(columns ...string) *Sql {
	s.query += ` (` + strings.Join(columns, ",") + `)`
	return s
}

func (s *Sql) Values(values ...any) *Sql {
	s.query += ` VALUES (`
	i := 0
	for ; i < len(values)-1; i++ {
		s.query += s.getNextParams() + `,`
		s.args = append(s.args, values[i])
	}
	s.query += s.getNextParams() + `)`
	s.args = append(s.args, values[i])
	return s
}

func (s *Sql) ColumnsValues(columnsValues []ColumnValue) *Sql {
	var columns []string
	var values []any
	for _, cv := range columnsValues {
		columns = append(columns, cv.Column)
		values = append(values, cv.Value)
	}
	s.query += ` (` + strings.Join(columns, ",") + `) VALUES (`
	i := 0
	for ; i < len(values)-1; i++ {
		s.query += s.getNextParams() + `,`
		s.args = append(s.args, values[i])
	}
	s.query += s.getNextParams() + `)`
	s.args = append(s.args, values[i])
	return s
}

func (s *Sql) Returning(columns ...string) *Sql {
	s.query += ` RETURNING ` + strings.Join(columns, ",")
	return s
}
