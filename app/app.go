package app

import (
	"io"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/saaremaa/work_calend/api"
	"github.com/saaremaa/work_calend/config"
	"github.com/saaremaa/work_calend/storage"
)

// App это базовая структура со всей необходимой информацией о приложении.
type App struct {
	StartTime time.Time
	Logger    *logrus.Logger
	Config    *config.Config
}

// New создаем новый экземпляр приложения
func New(cfg *config.Config, logger *logrus.Logger) (*App, error) {
	return &App{
		Logger:    logger,
		Config:    cfg,
		StartTime: time.Now(),
	}, nil
}

// Run запускаем приложение
func (a *App) Run() {
	st, err := storage.NewStorage(*a.Config, a.Logger)
	if err != nil && err != io.EOF {
		a.Logger.Fatalln("can't create new storage: ", err)
	}

	appi := api.NewAPI(st, a.Logger, a.Config)
	appi.InitRouter()
	api.ServeAPI(appi)
}
