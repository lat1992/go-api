package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (h *Handler) Register(c *gin.Context) {
	data := struct {
		Email       string `json:"email"`
		Password    string `json:"password"`
		CompanyName string `json:"company_name"`
		FullName    string `json:"full_name"`
		Country     string `json:"country"`
		Telephone   string `json:"telephone"`
	}{}
	if err := c.BindJSON(&data); err != nil {
		h.errorEndCall(c, http.StatusBadRequest, "FormatError", "User: BindJson: %v", err)
		return
	}
	if data.Email == "" || data.Password == "" || data.FullName == "" {
		h.warnEndCall(c, http.StatusForbidden, "EmailPasswordOrFullNameNotFound", "Auth: email/password/full name not found")
		return
	}
	err := h.controller.CreateUser(data.Email, data.Password, data.FullName)
	if err != nil {
		switch err.Error() {
		case "EmailUsed":
			h.infoEndCall(c, http.StatusForbidden, err.Error(), "Auth: email is used")
		default:
			h.errorEndCall(c, http.StatusBadRequest, "InternalError", "Auth: CreateUser: %v", err)
		}
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "created"})
}

func (h *Handler) Login(c *gin.Context) {
	data := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Token    string `json:"token"`
	}{}
	err := c.BindJSON(&data)
	if err != nil {
		h.errorEndCall(c, http.StatusBadRequest, "FormatError", "Auth: BindJson: %v", err)
		return
	}
	if data.Token == "" {
		h.warnEndCall(c, http.StatusUnauthorized, "NoToken", "Auth: token not found")
		return
	}
	if data.Email == "" || data.Password == "" {
		h.warnEndCall(c, http.StatusForbidden, "EmailOrPasswordNotFound", "Auth: email/password not found")
		return
	}
	accessToken, refreshToken, err := h.controller.LoginUser(data.Email, data.Password, data.Token)
	if err != nil {
		switch err.Error() {
		case "WrongCaptcha":
			h.warnEndCall(c, http.StatusForbidden, err.Error(), "Auth: wrong captcha token")
		case "EmailNotFound":
			h.infoEndCall(c, http.StatusForbidden, err.Error(), "Auth: email not found")
		case "PasswordIncorrect":
			h.infoEndCall(c, http.StatusForbidden, err.Error(), "Auth: password incorrect")
		default:
			h.errorEndCall(c, http.StatusBadRequest, "InternalError", "Auth: LoginUser: %v", err)
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"access_token": accessToken, "refresh_token": refreshToken})
}

func (h *Handler) AuthorizeJWT(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	claims, err := h.controller.CheckToken(authHeader)
	if err != nil {
		h.warnEndCall(c, http.StatusUnauthorized, "WrongAPIKey", "Auth: wrong api key: %v", err)
		return
	}
	if claims.ExpiresAt > time.Now().UTC().Unix() {
		c.Set("userId", claims.Id)
		c.Set("email", claims.Email)
		c.Next()
	} else {
		h.warnEndCall(c, http.StatusUnauthorized, "ExpiredAPIKey", "Auth: api key is expired")
		return
	}
}

func (h *Handler) RefreshToken(c *gin.Context) {
	refreshHeader := c.GetHeader("Authorization")
	accessToken, err := h.controller.LoginWithRefreshToken(refreshHeader)
	if err != nil {
		h.errorEndCall(c, http.StatusBadRequest, "InternalError", "Auth: RefreshToken: %v", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"access_token": accessToken, "refresh_token": refreshHeader})
}
