package query

func (s *Sql) Inner(tableB, columnA, columnB string) *Sql {
	s.query += ` INNER JOIN ` + tableB + ` ON ` + s.table + `.` + columnA + ` = ` + tableB + `.` + columnB
	return s
}

func (s *Sql) Left(tableB, columnA, columnB string) *Sql {
	s.query += ` LEFT JOIN ` + tableB + ` ON ` + s.table + `.` + columnA + ` = ` + tableB + `.` + columnB
	return s
}

func (s *Sql) LeftOuter(tableB, columnA, columnB string) *Sql {
	s.query += ` LEFT OUTER JOIN ` + tableB + ` ON ` + s.table + `.` + columnA + ` = ` + tableB + `.` + columnB
	return s
}

func (s *Sql) Right(tableB, columnA, columnB string) *Sql {
	s.query += ` RIGHT JOIN ` + tableB + ` ON ` + s.table + `.` + columnA + ` = ` + tableB + `.` + columnB
	return s
}

func (s *Sql) RightOuter(tableB, columnA, columnB string) *Sql {
	s.query += ` RIGHT OUTER JOIN ` + tableB + ` ON ` + s.table + `.` + columnA + ` = ` + tableB + `.` + columnB
	return s
}

func (s *Sql) Full(tableB, columnA, columnB string) *Sql {
	s.query += ` FULL JOIN ` + tableB + ` ON ` + s.table + `.` + columnA + ` = ` + tableB + `.` + columnB
	return s
}

func (s *Sql) FullOuter(tableB, columnA, columnB string) *Sql {
	s.query += ` FULL OUTER JOIN ` + tableB + ` ON ` + s.table + `.` + columnA + ` = ` + tableB + `.` + columnB
	return s
}
