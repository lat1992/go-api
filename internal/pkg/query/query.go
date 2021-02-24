package query

import (
	"strconv"
	"strings"
)

type sql struct {
	query string
}

func (s *sql) Build() string {
	return s.query
}

func (s *sql) countNextParams() string {
	return strconv.Itoa(strings.Count(s.query, "$") + 1)
}
