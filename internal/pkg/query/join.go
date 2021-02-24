package query

func (s *sql) Inner(table, columnA, columnB string) *sql {
	s.query += ` INNER JOIN ` + table + ` ON ` + columnA + `=` + columnB
	return s
}

func (s *sql) Left(table, columnA, columnB string) *sql {
	s.query += ` LEFT JOIN ` + table + ` ON ` + columnA + `=` + columnB
	return s
}

func (s *sql) LeftOuter(table, columnA, columnB string) *sql {
	s.query += ` LEFT OUTER JOIN ` + table + ` ON ` + columnA + `=` + columnB
	return s
}

func (s *sql) Right(table, columnA, columnB string) *sql {
	s.query += ` RIGHT JOIN ` + table + ` ON ` + columnA + `=` + columnB
	return s
}

func (s *sql) RightOuter(table, columnA, columnB string) *sql {
	s.query += ` RIGHT OUTER JOIN ` + table + ` ON ` + columnA + `=` + columnB
	return s
}

func (s *sql) Full(table, columnA, columnB string) *sql {
	s.query += ` FULL JOIN ` + table + ` ON ` + columnA + `=` + columnB
	return s
}

func (s *sql) FullOuter(table, columnA, columnB string) *sql {
	s.query += ` FULL OUTER JOIN ` + table + ` ON ` + columnA + `=` + columnB
	return s
}
