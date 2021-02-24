package query

import "strconv"

func Update(table string) *sql {
	return &sql{
		query: `UPDATE ` + table,
	}
}

func (s *sql) Set(columns ...string) *sql {
	s.query += ` SET `
	for i, column := range columns {
		if i < len(columns)-1 {
			s.query += column + ` = $` + strconv.Itoa(i+1) + `,`
		} else {
			s.query += column + ` = $` + strconv.Itoa(i+1)
		}
	}
	return s
}
