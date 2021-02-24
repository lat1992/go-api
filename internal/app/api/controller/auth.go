package controller

import (
	"fmt"
	"go-api/internal/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

func (c *Controller) CheckToken(authHeader string) (*jwt.Claims, error) {
	const bearerSchema = "Bearer "
	token := authHeader[len(bearerSchema):]
	return c.jwt.ParseToken(token, false)
}

func (c *Controller) LoginUser(email, password, captchaToken string) (string, string, error) {
	// WARN remove if you don't use reCaptcha
	check, err := c.captcha.Verify(captchaToken)
	if err != nil {
		return "", "", fmt.Errorf("CaptchaModuleError")
	}
	if !check {
		return "", "", fmt.Errorf("WrongCaptcha")
	}

	count, err := c.model.CountUserByEmail(email)
	if err != nil {
		return "", "", err
	}
	if count == 0 {
		return "", "", fmt.Errorf("EmailNotFound")
	}
	id, hash, err := c.model.SelectIdAndPasswordByEmail(email)
	if err != nil {
		return "", "", err
	}
	if bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) != nil {
		return "", "", fmt.Errorf("PasswordIncorrect")
	}
	accessToken, err := c.jwt.GenerateToken(id, email, false)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := c.jwt.GenerateToken(id, email, true)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, err
}

func (c *Controller) LoginWithRefreshToken(refreshToken string) (string, error) {
	const bearerSchema = "Bearer "
	token := refreshToken[len(bearerSchema):]
	claims, err := c.jwt.ParseToken(token, true)
	if err != nil {
		return "", err
	}
	accessToken, err := c.jwt.GenerateToken(claims.Id, claims.Email, false)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}
