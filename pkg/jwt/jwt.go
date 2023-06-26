package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type Service struct {
	accessSecret  string
	refreshSecret string
}

type Claims struct {
	Id    int64
	Email string
	jwt.StandardClaims
}

func NewService(accessSecret, refreshSecret string) *Service {
	return &Service{
		accessSecret:  accessSecret,
		refreshSecret: refreshSecret,
	}
}

const (
	accessExpiredTime  = 15 * time.Minute
	refreshExpiredTime = 15 * 24 * time.Hour
)

func (s *Service) GenerateToken(userId int64, email string, isRefresh bool) (string, error) {
	expired := time.Now().Add(accessExpiredTime).Unix()
	if isRefresh {
		expired = time.Now().Add(refreshExpiredTime).Unix()
	}
	claims := &Claims{
		Id:    userId,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expired,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if isRefresh {
		return token.SignedString([]byte(s.refreshSecret))
	}
	return token.SignedString([]byte(s.accessSecret))
}

func (s *Service) ParseToken(tokenString string, isRefresh bool) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		if isRefresh {
			return []byte(s.refreshSecret), nil
		}
		return []byte(s.accessSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims := token.Claims.(*Claims); token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("token not valid")
}
