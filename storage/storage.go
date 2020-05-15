package storage

import (
	"io"

	"github.com/saaremaa/work-calend/config"
	"github.com/saaremaa/work-calend/storage/csvstorage"
	"github.com/sirupsen/logrus"
)

type Storage struct {
	Logger *logrus.Logger
	DB     *csvstorage.CSVStorage
}

func NewStorage(cfg config.Config, logger *logrus.Logger) (*Storage, error) {
	res := &Storage{}
	db, err := csvstorage.NewCSV(cfg, logger)
	if err != nil && err != io.EOF {
		return res, err
	}

	res.DB = db
	res.Logger = logger
	return res, nil
}
