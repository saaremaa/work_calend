package storage

import (
	"io"

	"github.com/sirupsen/logrus"
	"github.com/saaremaa/work_calend/config"
	"github.com/saaremaa/work_calend/storage/csvstorage"
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
