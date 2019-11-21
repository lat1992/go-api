/*
 * File: X:\go-api\model\UserModel.go
 * Created At: 2019-11-05 22:56:10
 * Created By: Mauhoi WU
 * 
 * Modified At: 2019-11-21 19:12:22
 * Modified By: Mauhoi WU
 */

package model

import (
	"time"
)

type User struct {
	tableName	struct{}	`pg:"user"`
	Id			int			`pg:",pk"`
	Username	string		`pg:",unique,notnull"`
	Email		string		`pg:",unique,notnull"`
	Password	string
	FullName	string
	Country		string
	Telephone	string
	UpdateAt	time.Time
}

func VerifyUsername(username string) int {
	user := new(User)
	count, err := db.Model(user).Where("username LIKE ?", username).Count()
	if err != nil {
		panic(err)
	}
	return count
}

func VerifyEmail(email string) int {
	user := new(User)
	count, err := db.Model(user).Column("email").Where("email LIKE ?", email).Count()
	if err != nil {
		panic(err)
	}
	return count
}

func GetUserIdAndPasswordByUsername(username string) (int, string) {
	user := new(User)
	err := db.Model(user).Column("id", "password").Where("username LIKE ?", username).Select()
	if err != nil {
		panic(err)
	}
	return user.Id, user.Password
}

func VerifyUserIdAndPassword(id int, password string) int {
	user := new(User)
	count, err := db.Model(user).Where("id = ? AND password LIKE ?", id, password).Count()
	if err != nil {
		panic(err)
	}
	return count
}

func GetUserById(id int) *User {
	user := &User{Id: id}
	err := db.Select(user)
	if err != nil {
		panic(err)
	}
	return user
}

func GetUsers(limit int, offset int) []User {
	var users []User
	err := db.Model(&users).Limit(limit).Offset(offset).Select()
	if err != nil {
		panic(err)
	}
	return users
}

func AddUser(username, email, password, full_name, country, telephone string) int {
	user := User{
		Username: username,
		Email: email,
		Password: password,
		Country: country,
		FullName: full_name,
		Telephone: telephone,
		UpdateAt: time.Now(),
	}
	err := db.Insert(&user)
	if err != nil {
		panic(err)
	}
	return user.Id
}

func UpdateUser(id int, email, password, full_name, country, telephone string) {
	user := &User{
		Id: id,
		Email: email,
		Password: password,
		FullName: full_name,
		Country: country,
		Telephone: telephone,
	}
	err := db.Update(user)
	if err != nil {
		panic(err)
	}
}

func UpdateUserPassword(id int, password string) {
	user := &User{
		Id: id,
		Password: password,
	}
	err := db.Update(user)
	if err != nil {
		panic(err)
	}
}

func DeleteUser(id int) {
	user := &User{Id: id}
	err := db.Delete(user)
	if err != nil {
		panic(err)
	}
}
