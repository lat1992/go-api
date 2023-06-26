package database

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"go-api/internal"
	"go-api/pkg/query"
)

func (db *Database) CountUserByEmail(email string) (int, error) {
	var count int
	q := query.Select("count(1)").
		From(db.getTableName("user")).
		Where("email").Like(email)
	if err := db.postgres.QueryRow(context.Background(),
		q.Build(),
		q.Args()...).Scan(&count); err != nil {
		return 0, fmt.Errorf("CountUserByEmail: %w", err)
	}
	return count, nil
}

func (db *Database) SelectIdAndPasswordByEmail(email string) (int64, string, error) {
	var id int64
	var password string
	q := query.Select("id", "password").
		From(db.getTableName("user")).
		Where("email").Like(email)
	if err := db.postgres.QueryRow(context.Background(),
		q.Build(),
		q.Args()...).Scan(&id, &password); err != nil {
		return 0, "", fmt.Errorf("SelectIdAndPasswordByEmail: %w", err)
	}
	return id, password, nil
}

func (db *Database) SelectUserById(id int64) (internal.User, error) {
	q := query.Select("id", "email", "full_name", "update_at").
		From(db.getTableName("user")).
		Where("id").Equal(id)
	rows, err := db.postgres.Query(context.Background(),
		q.Build(),
		q.Args()...)
	if err != nil {
		return internal.User{}, fmt.Errorf("SelectUserById: %w", err)
	}
	var user internal.User
	if err = pgxscan.ScanOne(&user, rows); err != nil {
		return internal.User{}, fmt.Errorf("SelectUserById: %w", err)
	}
	return user, nil
}

func (db *Database) SelectUsers(limit int, offset int) ([]internal.User, error) {
	var users []internal.User
	rows, err := db.postgres.Query(context.Background(),
		query.Select("id", "email", "password", "full_name", "update_at").
			From(db.getTableName("user")).
			Limit(limit).
			Offset(offset).
			Build())
	if err != nil {
		return nil, fmt.Errorf("SelectUsers: %w", err)
	}
	if err = pgxscan.ScanAll(&users, rows); err != nil {
		return nil, fmt.Errorf("SelectUsers: %w", err)
	}
	return users, nil
}

func (db *Database) InsertUser(tx pgx.Tx, email, password, fullName string) (int64, error) {
	var id int64
	q := query.Insert(db.getTableName("user")).
		ColumnsValues([]query.ColumnValue{
			{Column: "email", Value: email},
			{Column: "password", Value: password},
			{Column: "full_name", Value: fullName},
		}).
		Returning("id")
	if err := tx.QueryRow(context.Background(),
		q.Build(),
		q.Args()...).Scan(&id); err != nil {
		return 0, fmt.Errorf("InsertUser: %w", err)
	}
	return id, nil
}

func (db *Database) UpdateUserById(tx pgx.Tx, id int64, email, password, fullName string) error {
	q := query.Update(db.getTableName("user")).
		Set([]query.ColumnValue{
			{Column: "email", Value: email},
			{Column: "password", Value: password},
			{Column: "full_name", Value: fullName},
		}).
		Where("id").Equal(id)
	if _, err := tx.Exec(context.Background(),
		q.Build(),
		q.Args()...); err != nil {
		return err
	}
	return nil
}

func (db *Database) UpdateUserPasswordById(tx pgx.Tx, id int64, password string) error {
	q := query.Update(db.getTableName("user")).
		Set([]query.ColumnValue{
			{Column: "password", Value: password},
		}).
		Where("id").Equal(id)
	if _, err := tx.Exec(context.Background(),
		q.Build(),
		q.Args()...); err != nil {
		return fmt.Errorf("UpdateUserPasswordById: %w", err)
	}
	return nil
}

func (db *Database) DeleteUser(tx pgx.Tx, id int64) error {
	q := query.Delete(db.getTableName("user")).
		Where("id").Equal(id)
	if _, err := tx.Exec(context.Background(),
		q.Build(),
		q.Args()...); err != nil {
		return fmt.Errorf("DeleteUser: %w", err)
	}
	return nil
}
