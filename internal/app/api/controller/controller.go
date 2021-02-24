package controller

import (
	"github.com/spf13/viper"
	"go-api/internal/app/api/model"
	"go-api/internal/pkg/captcha"
	"go-api/internal/pkg/jwt"
	"go.uber.org/zap"
)

type Controller struct {
	model   *model.Model
	captcha *captcha.Captcha
	jwt     *jwt.Service
	logger  *zap.SugaredLogger
	config  *viper.Viper
}

func NewController(conf *viper.Viper, logger *zap.SugaredLogger) *Controller {

	j := jwt.NewService(conf.GetString("jwt.accessSecret"), conf.GetString("jwt.refreshSecret"))
	capt := captcha.NewConnector(conf.GetString("captcha.secret")) // WARN remove if you don't use reCaptcha
	m := model.NewModel(logger, conf)
	return &Controller{
		model:   m,
		captcha: capt, // WARN remove if you don't use reCaptcha
		jwt:     j,
		logger:  logger,
		config:  conf,
	}
}

func (c *Controller) Destroy() {
	c.model.Close()
}
