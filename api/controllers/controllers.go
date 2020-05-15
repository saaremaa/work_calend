package controllers

import (
	"time"

	"github.com/labstack/echo"
	"github.com/saaremaa/work-calend/config"
	"github.com/saaremaa/work-calend/storage"
	"github.com/sirupsen/logrus"
)

type Controllers struct {
	StartTime  time.Time
	Logger     *logrus.Logger
	HttpPort   string
	Config     *config.Config
	Router     *echo.Echo
	Repository *storage.Storage
}

func NewControllers(rep *storage.Storage, log *logrus.Logger, cfg *config.Config) *Controllers {
	a := &Controllers{
		HttpPort:   cfg.HttpPort,
		StartTime:  time.Now(),
		Config:     cfg,
		Logger:     log,
		Repository: rep,
	}
	return a
}
