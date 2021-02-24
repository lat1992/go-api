package query

import (
	"strconv"
	"strings"
)

func Insert(table string) *sql {
	return &sql{
		query: `INSERT INTO ` + table,
	}
}

func (s *sql) Columns(columns ...string) *sql {
	s.query += ` (` + strings.Join(columns, ",") + `)`
	return s
}

func (s *sql) Values(count int) *sql {
	s.query += ` VALUES (`
	i := 1
	for i < count {
		s.query += `$` + strconv.Itoa(i) + `,`
		i++
	}
	s.query += `$` + strconv.Itoa(i) + `)`
	return s
}

func (s *sql) ColumnsValues(columns ...string) *sql {
	s.query += ` (` + strings.Join(columns, ",") + `) VALUES (`
	i := 1
	for i < len(columns) {
		s.query += `$` + strconv.Itoa(i) + `,`
		i++
	}
	s.query += `$` + strconv.Itoa(i) + `)`
	return s
}

func (s *sql) Returning(columns ...string) *sql {
	s.query += ` RETURNING ` + strings.Join(columns, ",")
	return s
}
