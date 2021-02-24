package controller

import (
	"fmt"
	"go-api/internal/app/api/model"
	"golang.org/x/crypto/bcrypt"
)

func (c *Controller) CreateUser(email, password, fullName string) error {
	count, err := c.model.CountUserByEmail(email)
	if err != nil {
		return err
	}
	if count != 0 {
		return fmt.Errorf("EmailUsed")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	tx, err := c.model.BeginTransaction()
	if err != nil {
		return err
	}
	_, err = c.model.InsertUser(tx, email, string(hash), fullName)
	if err != nil {
		if err := c.model.RollbackTransaction(tx); err != nil {
			return err
		}
		return err
	}
	if err := c.model.CommitTransaction(tx); err != nil {
		if err := c.model.RollbackTransaction(tx); err != nil {
			return err
		}
		return err
	}
	return nil
}

func (c *Controller) GetUsers(rows, page int) ([]model.User, error) {
	return c.model.SelectUsers(rows, page)
}

func (c *Controller) GetUser(id int64) (model.User, error) {
	return c.model.SelectUserById(id)
}

func (c *Controller) ModifyUserPassword(id int64, password string) error {
	tx, err := c.model.BeginTransaction()
	if err != nil {
		return err
	}
	if err := c.model.UpdateUserPasswordById(tx, id, password); err != nil {
		if err := c.model.RollbackTransaction(tx); err != nil {
			return err
		}
		return err
	}
	if err := c.model.CommitTransaction(tx); err != nil {
		return err
	}
	return nil
}

func (c *Controller) ModifyUser(id int64, email, password, fullName string) error {
	tx, err := c.model.BeginTransaction()
	if err != nil {
		return err
	}
	if err := c.model.UpdateUserById(tx, id, email, password, fullName); err != nil {
		if err := c.model.RollbackTransaction(tx); err != nil {
			return err
		}
		return err
	}
	if err := c.model.CommitTransaction(tx); err != nil {
		return err
	}
	return nil
}

func (c *Controller) DeleteUser(id int64) error {
	tx, err := c.model.BeginTransaction()
	if err != nil {
		return err
	}
	if err := c.model.DeleteUser(tx, id); err != nil {
		if err := c.model.RollbackTransaction(tx); err != nil {
			return err
		}
		return err
	}
	if err := c.model.CommitTransaction(tx); err != nil {
		return err
	}
	return nil
}
