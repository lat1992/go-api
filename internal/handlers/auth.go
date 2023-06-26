package handlers

import (
	"github.com/gin-gonic/gin"
	"go-api/internal"
	"net/http"
	"time"
)

type registerRequest struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	CompanyName string `json:"company_name"`
	FullName    string `json:"full_name"`
	Country     string `json:"country"`
	Telephone   string `json:"telephone"`
}

func Register(us internal.UserService) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data registerRequest
		if err := c.BindJSON(&data); err != nil {
			warnEndCall(c, http.StatusBadRequest, "FormatError", "User: BindJson", err)
			return
		}
		if data.Email == "" || data.Password == "" || data.FullName == "" {
			debugEndCall(c, http.StatusForbidden, "EmailPasswordOrFullNameNotFound", "Auth: email/password/full name not found")
			return
		}
		err := us.CreateUser(data.Email, data.Password, data.FullName)
		if err != nil {
			switch err {
			case internal.ErrEmailExist:
				debugEndCall(c, http.StatusForbidden, err.Error(), "Auth: email is used")
			default:
				errorEndCall(c, http.StatusBadRequest, "InternalError", "Auth: CreateUser", err)
			}
			return
		}
		c.JSON(http.StatusCreated, gin.H{"status": "created"})
	}
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type tokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func Login(as internal.AuthService) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data loginRequest
		if err := c.BindJSON(&data); err != nil {
			warnEndCall(c, http.StatusBadRequest, "FormatError", "Auth", err)
			return
		}
		if data.Token == "" {
			infoEndCall(c, http.StatusUnauthorized, "NoToken", "Auth: token not found")
			return
		}
		if data.Email == "" || data.Password == "" {
			debugEndCall(c, http.StatusForbidden, "EmailOrPasswordNotFound", "Auth: email/password not found")
			return
		}
		accessToken, refreshToken, err := as.LoginUser(data.Email, data.Password, data.Token)
		if err != nil {
			switch err {
			case internal.ErrWrongCaptcha:
				warnEndCall(c, http.StatusForbidden, err.Error(), "Auth: wrong captcha token")
			case internal.ErrEmailNotFound:
				debugEndCall(c, http.StatusForbidden, err.Error(), "Auth: email not found")
			case internal.ErrPasswordIncorrect:
				infoEndCall(c, http.StatusForbidden, err.Error(), "Auth: password incorrect")
			default:
				errorEndCall(c, http.StatusBadRequest, "InternalError", "Auth", err)
			}
			return
		}
		c.JSON(http.StatusOK, tokenResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		})
	}
}

func AuthorizeJWT(as internal.AuthService) func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		claims, err := as.CheckToken(authHeader)
		if err != nil {
			warnEndCall(c, http.StatusUnauthorized, "WrongAPIKey", "Auth: wrong api key", err)
			return
		}
		if claims.ExpiresAt > time.Now().UTC().Unix() {
			c.Set("userId", claims.Id)
			c.Set("email", claims.Email)
			c.Next()
		} else {
			debugEndCall(c, http.StatusUnauthorized, "ExpiredAPIKey", "Auth: api key is expired")
			return
		}
	}
}

func RefreshToken(as internal.AuthService) func(c *gin.Context) {
	return func(c *gin.Context) {
		refreshHeader := c.GetHeader("Authorization")
		accessToken, err := as.LoginWithRefreshToken(refreshHeader)
		if err != nil {
			errorEndCall(c, http.StatusBadRequest, "InternalError", "Auth: RefreshToken", err)
			return
		}
		c.JSON(http.StatusOK, tokenResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshHeader,
		})
	}
}
