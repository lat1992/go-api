package services

import (
	"fmt"
	"go-api/internal"
	"go-api/internal/database"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	database *database.Database
}

func NewUserService(db *database.Database) *UserService {
	return &UserService{
		database: db,
	}
}

func (s *UserService) CreateUser(email, password, fullName string) error {
	count, err := s.database.CountUserByEmail(email)
	if err != nil {
		return fmt.Errorf("CreateUser: %w", err)
	}
	if count != 0 {
		return internal.ErrEmailExist
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return fmt.Errorf("CreateUser: %w", err)
	}
	tx, err := s.database.BeginTransaction()
	if err != nil {
		return fmt.Errorf("CreateUser: %w", err)
	}
	_, err = s.database.InsertUser(tx, email, string(hash), fullName)
	if err != nil {
		if err = s.database.RollbackTransaction(tx); err != nil {
			return fmt.Errorf("CreateUser: %w", err)
		}
		return fmt.Errorf("CreateUser: %w", err)
	}
	if err = s.database.CommitTransaction(tx); err != nil {
		if err = s.database.RollbackTransaction(tx); err != nil {
			return fmt.Errorf("CreateUser: %w", err)
		}
		return fmt.Errorf("CreateUser: %w", err)
	}
	return nil
}

func (s *UserService) GetUsers(rows, page int) ([]internal.User, error) {
	return s.database.SelectUsers(rows, page)
}

func (s *UserService) GetUser(id int64) (internal.User, error) {
	return s.database.SelectUserById(id)
}

func (s *UserService) ModifyUserPassword(id int64, password string) error {
	tx, err := s.database.BeginTransaction()
	if err != nil {
		return fmt.Errorf("ModifyUserPassword: %w", err)
	}
	if err = s.database.UpdateUserPasswordById(tx, id, password); err != nil {
		if err = s.database.RollbackTransaction(tx); err != nil {
			return fmt.Errorf("ModifyUserPassword: %w", err)
		}
		return fmt.Errorf("ModifyUserPassword: %w", err)
	}
	if err = s.database.CommitTransaction(tx); err != nil {
		return fmt.Errorf("ModifyUserPassword: %w", err)
	}
	return nil
}

func (s *UserService) ModifyUser(id int64, email, password, fullName string) error {
	tx, err := s.database.BeginTransaction()
	if err != nil {
		return fmt.Errorf("ModifyUser: %w", err)
	}
	if err = s.database.UpdateUserById(tx, id, email, password, fullName); err != nil {
		if err = s.database.RollbackTransaction(tx); err != nil {
			return fmt.Errorf("ModifyUser: %w", err)
		}
		return fmt.Errorf("ModifyUser: %w", err)
	}
	if err = s.database.CommitTransaction(tx); err != nil {
		return fmt.Errorf("ModifyUser: %w", err)
	}
	return nil
}

func (s *UserService) DeleteUser(id int64) error {
	tx, err := s.database.BeginTransaction()
	if err != nil {
		return fmt.Errorf("DeleteUser: %w", err)
	}
	if err = s.database.DeleteUser(tx, id); err != nil {
		if err = s.database.RollbackTransaction(tx); err != nil {
			return fmt.Errorf("DeleteUser: %w", err)
		}
		return fmt.Errorf("DeleteUser: %w", err)
	}
	if err = s.database.CommitTransaction(tx); err != nil {
		return fmt.Errorf("DeleteUser: %w", err)
	}
	return nil
}
