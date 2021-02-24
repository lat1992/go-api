package query

import (
	"strconv"
	"strings"
)

func Select(columns ...string) *sql {
	return &sql{
		query: `SELECT ` + strings.Join(columns, ","),
	}
}

func (s *sql) From(table string) *sql {
	s.query += ` FROM ` + table
	return s
}

func (s *sql) Limit(limit int) *sql {
	s.query += ` LIMIT ` + strconv.Itoa(limit)
	return s
}

func (s *sql) Offset(offset int) *sql {
	s.query += ` OFFSET ` + strconv.Itoa(offset)
	return s
}
