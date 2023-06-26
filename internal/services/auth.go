package services

import (
	"fmt"
	"go-api/internal"
	"go-api/internal/database"
	"go-api/pkg/jwt"
	"go-api/pkg/recaptcha"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	database *database.Database
	jwt      *jwt.Service
	captcha  *recaptcha.Captcha
}

func NewAuthService(db *database.Database, j *jwt.Service, c *recaptcha.Captcha) *AuthService {
	return &AuthService{
		database: db,
		jwt:      j,
		captcha:  c,
	}
}

func (s *AuthService) CheckToken(authHeader string) (*jwt.Claims, error) {
	const bearerSchema = "Bearer "
	token := authHeader[len(bearerSchema):]
	return s.jwt.ParseToken(token, false)
}

func (s *AuthService) LoginUser(email, password, captchaToken string) (string, string, error) {
	check, err := s.captcha.Verify(captchaToken)
	if err != nil {
		return "", "", fmt.Errorf("LoginUser: %w", err)
	}
	if !check {
		return "", "", internal.ErrWrongCaptcha
	}

	count, err := s.database.CountUserByEmail(email)
	if err != nil {
		return "", "", fmt.Errorf("LoginUser: %w", err)
	}
	if count == 0 {
		return "", "", internal.ErrEmailNotFound
	}
	id, hash, err := s.database.SelectIdAndPasswordByEmail(email)
	if err != nil {
		return "", "", fmt.Errorf("LoginUser: %w", err)
	}
	if bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) != nil {
		return "", "", internal.ErrPasswordIncorrect
	}
	accessToken, err := s.jwt.GenerateToken(id, email, false)
	if err != nil {
		return "", "", fmt.Errorf("LoginUser: %w", err)
	}
	refreshToken, err := s.jwt.GenerateToken(id, email, true)
	if err != nil {
		return "", "", fmt.Errorf("LoginUser: %w", err)
	}
	return accessToken, refreshToken, fmt.Errorf("LoginUser: %w", err)
}

func (s *AuthService) LoginWithRefreshToken(refreshToken string) (string, error) {
	const bearerSchema = "Bearer "
	token := refreshToken[len(bearerSchema):]
	claims, err := s.jwt.ParseToken(token, true)
	if err != nil {
		return "", err
	}
	accessToken, err := s.jwt.GenerateToken(claims.Id, claims.Email, false)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}
