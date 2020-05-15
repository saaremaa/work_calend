package api

import (
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type (
	Stats struct {
		Uptime       time.Time      `json:"uptime"`
		RequestCount uint64         `json:"requestCount"`
		Statuses     map[string]int `json:"statuses"`
		AVGLatency   float64        `json:"avg_latency"`
		mutex        sync.RWMutex
	}
)

func NewStats() *Stats {
	return &Stats{
		Uptime:   time.Now(),
		Statuses: map[string]int{},
	}
}

// Process is the middleware function.
func (s *Stats) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			c.Error(err)
		}
		s.mutex.Lock()
		defer s.mutex.Unlock()
		s.RequestCount++
		status := strconv.Itoa(c.Response().Status)
		//latency = c.Response().Header.
		s.Statuses[status]++
		return nil
	}
}

// Handle is the endpoint to get stats.
func (s *Stats) Handle(c echo.Context) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return c.JSON(http.StatusOK, s)
}

// ServerHeader middleware adds a `Server` header to the response.
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
		return next(c)
	}
}

func (a *API) InitRouter() {
	s := NewStats()
	e := echo.New()
	e.Server.Addr = ":" + a.Config.HttpPort

	// Echo middleware init
	e.Use(s.Process)
	e.Use(middleware.BodyLimit("1M"))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Output: a.Logger.Out}))
	e.Use(middleware.Recover())

	// Static file
	e.File("/favicon.ico", "images/favicon.ico")

	// Маршруты
	e.GET("/", a.RespROOT)                       // * стартовая страница при обращении выкидывает на сайт компании
	e.GET("/api/v1/health", a.Health)            // * проверка доступности API для Docker
	e.GET("/api/v1/stat", a.Stat)                // * данные по датасету
	e.GET("/api/v1/stats", s.Handle)             // Endpoint to get stats
	e.GET("/api/v1/check_day/:date", a.CheckDay) // * проверка является ли день выходным

	a.Router = e
}
