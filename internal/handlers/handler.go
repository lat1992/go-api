package handlers

import (
	"github.com/gin-gonic/gin"
	"go-api/internal"
	"golang.org/x/exp/slog"
)

func GetRouter(as internal.AuthService, us internal.UserService) *gin.Engine {
	router := gin.Default()

	router.GET("/", Health)
	router.GET("/health", Health)

	auth := router.Group("/auth")
	auth.POST("/register", Register(us))
	auth.POST("/login", Login(as))
	auth.POST("/refresh", RefreshToken(as))
	auth.GET("/me", AuthorizeJWT(as), GetUser(us))

	user := router.Group("/user", AuthorizeJWT(as))
	user.GET("/:id", GetUser(us))
	user.DELETE("/delete/", DeleteUser(us))

	users := router.Group("/users", AuthorizeJWT(as))
	users.GET("/:rows/:page", GetUsers(us))

	return router
}

func debugEndCall(c *gin.Context, code int, message string, logMessage string, args ...any) {
	if gin.IsDebugging() {
		slog.Debug(logMessage, args)
	}
	c.JSON(code, gin.H{"message": message})
	c.Abort()
}

func infoEndCall(c *gin.Context, code int, message string, logMessage string, args ...any) {
	slog.Info(logMessage, args)
	c.JSON(code, gin.H{"message": message})
	c.Abort()
}

func warnEndCall(c *gin.Context, code int, message string, logMessage string, args ...any) {
	slog.Warn(logMessage, args)
	c.JSON(code, gin.H{"message": message})
	c.Abort()
}

func errorEndCall(c *gin.Context, code int, message string, logMessage string, args ...any) {
	slog.Error(logMessage, args)
	c.JSON(code, gin.H{"message": message})
	c.Abort()
}
