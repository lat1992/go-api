package service

import (
	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-api/internal/app/api/handler"
	"go.uber.org/zap"
	"time"
)

type Service struct {
	router  *gin.Engine
	handler *handler.Handler
	logger  *zap.SugaredLogger
	config  *viper.Viper
}

func NewService(conf *viper.Viper, logger *zap.SugaredLogger) *Service {
	gin.SetMode(conf.GetString("app.mode"))
	router := gin.New()
	if gin.IsDebugging() {
		router.Use(ginzap.Ginzap(logger.Desugar(), time.RFC3339, true))
	}
	router.Use(ginzap.RecoveryWithZap(logger.Desugar(), true))
	h := handler.NewHandler(conf, logger)
	return &Service{
		router:  router,
		handler: h,
		logger:  logger,
		config:  conf,
	}
}

func (s *Service) Destroy() {
	s.handler.Destroy()
}

func (s *Service) Start() {
	s.SetRouter()
	if s.config.GetBool("app.tls.enable") {
		if err := autotls.Run(s.router, s.config.GetString("app.domain")); err != nil {
			s.logger.Fatalf("Service start: autoTls run: %v", err)
		}
	} else {
		if err := s.router.Run(":" + s.config.GetString("app.port")); err != nil {
			s.logger.Fatalf("Service start: router run: %v", err)
		}
	}
}
