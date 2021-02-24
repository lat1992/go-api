package main

import (
	"go-api/internal/app/api/config"
	"go-api/internal/app/api/service"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

func main() {
	conf := config.GetConfiguration()
	logger, err := getZapLogger(conf.GetString("app.env"))
	if err != nil {
		log.Fatalf("Logger initialization: %v", err)
	}
	s := service.NewService(conf, logger)
	defer s.Destroy()
	s.Start()
}

func getZapLogger(env string) (*zap.SugaredLogger, error) {
	var conf zap.Config
	var level zapcore.Level
	if env == "production" {
		conf = zap.NewProductionConfig()
		if err := level.Set("warn"); err != nil {
			return nil, err
		}
	} else if env == "staging" {
		conf = zap.NewProductionConfig()
		if err := level.Set("info"); err != nil {
			return nil, err
		}
	} else {
		conf = zap.NewDevelopmentConfig()
		if err := level.Set("debug"); err != nil {
			return nil, err
		}
	}
	zapLogger, err := conf.Build()
	if err != nil {
		return nil, err
	}
	return zapLogger.Sugar(), nil
}
