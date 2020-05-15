package api

import (
	"time"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/tylerb/graceful"
	"github.com/saaremaa/work_calend/api/controllers"
	"github.com/saaremaa/work_calend/config"
	"github.com/saaremaa/work_calend/storage"
)

type API struct {
	*controllers.Controllers
	StartTime  time.Time
	Logger     *logrus.Logger
	HttpPort   string
	Config     *config.Config
	Router     *echo.Echo
	Repository *storage.Storage
}

func NewAPI(rep *storage.Storage, log *logrus.Logger, cfg *config.Config) *API {
	a := &API{
		controllers.NewControllers(rep, log, cfg),
		time.Now(),
		log,
		cfg.HttpPort,
		cfg,
		echo.New(),
		rep,
	}
	a.InitRouter()
	return a
}

func ServeAPI(api *API) {

	err := graceful.ListenAndServe(api.Router.Server, 5*time.Second)
	if err != nil {
		api.Logger.Error(err)
	}

}
