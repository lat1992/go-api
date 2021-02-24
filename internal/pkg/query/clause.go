package query

func (s *sql) Where(condition string) *sql {
	s.query += ` WHERE ` + condition + ` $` + s.countNextParams()
	return s
}

func (s *sql) Or(condition string) *sql {
	s.query += ` OR ` + condition + ` $` + s.countNextParams()
	return s
}

func (s *sql) And(condition string) *sql {
	s.query += ` AND ` + condition + ` $` + s.countNextParams()
	return s
}

func (s *sql) WhereEqual(column string) *sql {
	s.query += ` WHERE ` + column + ` = $` + s.countNextParams()
	return s
}

func (s *sql) OrEqual(column string) *sql {
	s.query += ` OR ` + column + ` = $` + s.countNextParams()
	return s
}

func (s *sql) AndEqual(column string) *sql {
	s.query += ` AND ` + column + ` = $` + s.countNextParams()
	return s
}

func (s *sql) WhereLike(column string) *sql {
	s.query += ` WHERE ` + column + ` LIKE $` + s.countNextParams()
	return s
}

func (s *sql) OrLike(column string) *sql {
	s.query += ` OR ` + column + ` LIKE $` + s.countNextParams()
	return s
}

func (s *sql) AndLike(column string) *sql {
	s.query += ` AND ` + column + ` LIKE $` + s.countNextParams()
	return s
}
