package model

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"go-api/internal/pkg/query"
	"time"
)

type User struct {
	Id        int64     `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	FullName  string    `json:"full_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
}

func (m *Model) CountUserByEmail(email string) (int, error) {
	var count int
	if err := m.database.postgres.QueryRow(context.Background(),
		query.Select("count(1)").
			From(m.database.getTableName("user")).
			WhereLike("email").
			Build(),
		email).Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

func (m *Model) SelectIdAndPasswordByEmail(email string) (int64, string, error) {
	var id int64
	var password string
	if err := m.database.postgres.QueryRow(context.Background(),
		query.Select("id", "password").
			From(m.database.getTableName("user")).
			WhereLike("email").
			Build(),
		email).Scan(&id, &password); err != nil {
		return 0, "", err
	}
	return id, password, nil
}

func (m *Model) SelectUserById(id int64) (User, error) {
	rows, err := m.database.postgres.Query(context.Background(),
		query.Select("id", "email", "full_name", "update_at").
			From(m.database.getTableName("user")).
			WhereEqual("id").
			Build(), id)
	if err != nil {
		return User{}, err
	}
	var user User
	if err := pgxscan.ScanOne(&user, rows); err != nil {
		return User{}, err
	}
	return user, nil
}

func (m *Model) SelectUsers(limit int, offset int) ([]User, error) {
	var users []User
	rows, err := m.database.postgres.Query(context.Background(),
		query.Select("id", "email", "password", "full_name", "update_at").
			From(m.database.getTableName("user")).
			Limit(limit).
			Offset(offset).
			Build())
	if err != nil {
		return nil, err
	}
	if err := pgxscan.ScanAll(&users, rows); err != nil {
		return nil, err
	}
	return users, nil
}

func (m *Model) InsertUser(tx pgx.Tx, email, password, fullName string) (int64, error) {
	var id int64
	if err := tx.QueryRow(context.Background(),
		query.Insert(m.database.getTableName("user")).
			ColumnsValues("email", "password", "full_name").
			Returning("id").
			Build(),
		email, password, fullName).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (m *Model) UpdateUserById(tx pgx.Tx, id int64, email, password, fullName string) error {
	if _, err := tx.Exec(context.Background(),
		query.Update(m.database.getTableName("user")).
			Set("email", "password", "full_name").
			WhereEqual("id").
			Build(),
		email, password, fullName, id); err != nil {
		return err
	}
	return nil
}

func (m *Model) UpdateUserPasswordById(tx pgx.Tx, id int64, password string) error {
	if _, err := tx.Exec(context.Background(),
		query.Update(m.database.getTableName("user")).
			Set("password").
			WhereEqual("id").
			Build(),
		password, id); err != nil {
		return err
	}
	return nil
}

func (m *Model) DeleteUser(tx pgx.Tx, id int64) error {
	if _, err := tx.Exec(context.Background(),
		query.Delete(m.database.getTableName("user")).
			WhereEqual("id").
			Build(),
		id); err != nil {
		return err
	}
	return nil
}
