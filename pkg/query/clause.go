package query

func (s *Sql) Where(column string) *Sql {
	s.query += ` WHERE ` + column
	return s
}

func (s *Sql) Or(column string) *Sql {
	s.query += ` OR ` + column
	return s
}

func (s *Sql) And(column string) *Sql {
	s.query += ` AND ` + column
	return s
}

func (s *Sql) Having(column string) *Sql {
	s.query += ` HAVING ` + column
	return s
}

func (s *Sql) Is(value any) *Sql {
	s.query += ` IS ` + s.getNextParams()
	s.args = append(s.args, value)
	return s
}

func (s *Sql) IsNull() *Sql {
	s.query += ` IS NULL`
	return s
}

func (s *Sql) IsNotNull() *Sql {
	s.query += ` IS NOT NULL`
	return s
}

func (s *Sql) Like(value any) *Sql {
	s.query += ` LIKE ` + s.getNextParams()
	s.args = append(s.args, value)
	return s
}

func (s *Sql) NotLike(value any) *Sql {
	s.query += ` NOT LIKE ` + s.getNextParams()
	s.args = append(s.args, value)
	return s
}

func (s *Sql) In(values ...any) *Sql {
	s.query += ` IN (`
	for i := 1; i < len(values); i++ {
		s.query += `$` + s.getNextParams() + `,`
	}
	s.query += `$` + s.getNextParams() + `)`
	s.args = append(s.args, values...)
	return s
}

func (s *Sql) Equal(value any) *Sql {
	s.query += ` = ` + s.getNextParams()
	s.args = append(s.args, value)
	return s
}

func (s *Sql) NotEqual(value any) *Sql {
	s.query += ` != ` + s.getNextParams()
	s.args = append(s.args, value)
	return s
}

func (s *Sql) GreaterThan(value any) *Sql {
	s.query += ` > ` + s.getNextParams()
	s.args = append(s.args, value)
	return s
}

func (s *Sql) GreaterThanOrEqual(value any) *Sql {
	s.query += ` >= ` + s.getNextParams()
	s.args = append(s.args, value)
	return s
}

func (s *Sql) LessThan(value any) *Sql {
	s.query += ` < ` + s.getNextParams()
	s.args = append(s.args, value)
	return s
}

func (s *Sql) LessThanOrEqual(value any) *Sql {
	s.query += ` <= ` + s.getNextParams()
	s.args = append(s.args, value)
	return s
}

func (s *Sql) Between(value1, value2 any) *Sql {
	s.query += ` BETWEEN ` + s.getNextParams() + ` AND ` + s.getNextParams()
	s.args = append(s.args, value1, value2)
	return s
}

func (s *Sql) NotBetween(value1, value2 any) *Sql {
	s.query += ` NOT BETWEEN ` + s.getNextParams() + ` AND ` + s.getNextParams()
	s.args = append(s.args, value1, value2)
	return s
}

func (s *Sql) Exists(value any) *Sql {
	s.query += ` EXISTS ` + s.getNextParams()
	s.args = append(s.args, value)
	return s
}

func (s *Sql) NotExists(value any) *Sql {
	s.query += ` NOT EXISTS ` + s.getNextParams()
	s.args = append(s.args, value)
	return s
}
