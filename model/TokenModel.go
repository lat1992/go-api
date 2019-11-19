/*
 * File: X:\go-api\model\TokenModel.go
 * Created At: 2019-11-06 14:57:07
 * Created By: Mauhoi WU
 * 
 * Modified At: 2019-11-19 17:19:08
 * Modified By: Mauhoi WU
 */

package model

import (
	"time"
)

type Token struct {
	tableName	struct{}	`pg:"token"`
	Id			int			`pg:",pk"`
	UserId		int			`pg:",unique,notnull"`
	Token		string		`pg:",unique"`
	ExpiredAt	time.Time
}

func (m *Model) CreateToken(user_id int, generated_token string, expired_time time.Time) {
	token := Token{
		UserId: user_id,
		Token: generated_token,
		ExpiredAt: expired_time,
	}
	err := m.db.Insert(&token)
	if err != nil {
		panic(err)
	}
}

func (m *Model) VerifyToken(token_string string) int {
	token := new(Token)
	count, err := m.db.Model(token).Where("token LIKE ?", token_string).Count()
	if err != nil {
		panic(err)
	}
	return count
}

func (m *Model) GetUserIdByToken(token_string string) int {
	token := new(Token)
	err := m.db.Model(token).Column("user_id").Where("token LIKE ?", token_string).Select()
	if err != nil {
		panic(err)
	}
	return token.UserId
}

func (m *Model) DeleteOldToken(user_id int) {
	token := Token{ UserId: user_id }
	_, err := m.db.Model(&token).WherePK().Delete()
	if err != nil {
		panic(err)
	}
}
