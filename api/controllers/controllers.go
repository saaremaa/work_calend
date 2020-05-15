package controllers

import (
	"time"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/saaremaa/work_calend/config"
	"github.com/saaremaa/work_calend/storage"
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
