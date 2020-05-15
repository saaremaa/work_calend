package main

import (
	"path/filepath"
	"runtime"

	"github.com/saaremaa/work-calend/app"
	"github.com/saaremaa/work-calend/config"
	"github.com/sirupsen/logrus"
)

func main() {
	lg := logrus.New()

	cfg := config.NewConfig(getStartDir())

	a, err := app.New(cfg, lg)
	if err != nil {
		logrus.Fatal(err)
	}

	a.Run()
}

func getStartDir() string {
	_, fileName, _, _ := runtime.Caller(0)
	prefixPath := filepath.Dir(fileName)
	return prefixPath
}
