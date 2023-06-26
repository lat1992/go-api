package internal

import (
	"github.com/lat1992/go-api/pkg/jwt"
)

type AuthService interface {
	CheckToken(authHeader string) (*jwt.Claims, error)

	LoginUser(email, password, captchaToken string) (string, string, error)

	LoginWithRefreshToken(refreshToken string) (string, error)
}

type UserService interface {
	CreateUser(email, password, fullName string) error

	GetUsers(rows, page int) ([]User, error)

	GetUser(id int64) (User, error)

	ModifyUserPassword(id int64, password string) error

	ModifyUser(id int64, email, password, fullName string) error

	DeleteUser(id int64) error
}
