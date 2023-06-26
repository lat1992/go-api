package query

import (
	"strconv"
	"strings"
)

func Select(columns ...string) *Sql {
	return &Sql{
		query: `SELECT ` + strings.Join(columns, ","),
	}
}

func (s *Sql) From(table string) *Sql {
	s.query += ` FROM ` + table
	s.table = table
	return s
}

func (s *Sql) Limit(limit int) *Sql {
	s.query += ` LIMIT ` + strconv.Itoa(limit)
	return s
}

func (s *Sql) Offset(offset int) *Sql {
	s.query += ` OFFSET ` + strconv.Itoa(offset)
	return s
}

func (s *Sql) GroupBy(columns ...string) *Sql {
	s.query += ` GROUP BY ` + strings.Join(columns, ",")
	return s
}
