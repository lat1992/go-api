package main

import (
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"go-api/internal/config"
	"go-api/internal/database"
	"go-api/internal/handlers"
	"go-api/internal/services"
	"go-api/pkg/jwt"
	"go-api/pkg/recaptcha"
	"golang.org/x/exp/slog"
)

func main() {
	conf := config.GetConfiguration()
	db, err := database.NewDatabase(conf.GetString("postgres.host"), conf.GetString("postgres.port"), conf.GetString("postgres.database"), conf.GetString("postgres.user"), conf.GetString("postgres.password"), conf.GetString("postgres.schema"))
	if err != nil {
		slog.Error("Database connection error", err)
	}
	j := jwt.NewService(conf.GetString("jwt.secret"), conf.GetString("jwt.refresh_secret"))
	c := recaptcha.NewConnector(conf.GetString("recaptcha.secret"))

	as := services.NewAuthService(db, j, c)
	us := services.NewUserService(db)

	r := handlers.GetRouter(as, us)

	start(r, conf.GetString("app.domain"), conf.GetString("app.port"), conf.GetBool("app.tls.enable"))
}

func start(r *gin.Engine, domain, port string, tlsEnabled bool) {
	if tlsEnabled {
		if err := autotls.Run(r, domain); err != nil {
			slog.Error("Service start: autoTls run", err)
		}
	} else {
		if err := r.Run(":" + port); err != nil {
			slog.Error("Service start: router run", err)
		}
	}
}
