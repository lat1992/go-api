package query

import (
	"strconv"
)

type Sql struct {
	query      string
	table      string
	args       []any
	paramCount int
}

type ColumnValue struct {
	Column string
	Value  any
}

func (s *Sql) Build() string {
	return s.query
}

func (s *Sql) Args() []any {
	return s.args
}

func (s *Sql) getNextParams() string {
	s.paramCount++
	return `$` + strconv.Itoa(s.paramCount)
}
