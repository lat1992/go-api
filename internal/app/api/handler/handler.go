package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-api/internal/app/api/controller"
	"go.uber.org/zap"
)

type Handler struct {
	controller *controller.Controller
	logger     *zap.SugaredLogger
}

func NewHandler(conf *viper.Viper, logger *zap.SugaredLogger) *Handler {
	c := controller.NewController(conf, logger)
	return &Handler{
		controller: c,
		logger:     logger,
	}
}

func (h *Handler) Destroy() {
	h.controller.Destroy()
}

func (h *Handler) infoEndCall(c *gin.Context, code int, message string, template string, args ...interface{}) {
	if gin.IsDebugging() {
		if len(args) > 0 {
			h.logger.Infof(template, args)
		} else {
			h.logger.Infof(template)
		}
	}
	c.JSON(code, gin.H{"message": message})
	c.Abort()
}

func (h *Handler) warnEndCall(c *gin.Context, code int, message string, template string, args ...interface{}) {
	if len(args) > 0 {
		h.logger.Warnf(template, args)
	} else {
		h.logger.Warnf(template)
	}
	c.JSON(code, gin.H{"message": message})
	c.Abort()
}

func (h *Handler) errorEndCall(c *gin.Context, code int, message string, template string, args ...interface{}) {
	if len(args) > 0 {
		h.logger.Errorf(template, args)
	} else {
		h.logger.Errorf(template)
	}
	c.JSON(code, gin.H{"message": message})
	c.Abort()
}
