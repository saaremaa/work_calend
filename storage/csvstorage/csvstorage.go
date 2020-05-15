package csvstorage

import (
	"encoding/csv"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/saaremaa/work_calend/config"
)

// Storage структура описания хранилища
type CSVStorage struct {
	Logger *logrus.Logger
	DB     *map[string]map[string][]string
}

// NewCSV - создаем новое хранилище из CSV файла и заполняем его данными
func NewCSV(cfg config.Config, logger *logrus.Logger) (St *CSVStorage, err error) {
	St = new(CSVStorage)
	f, err := os.Open(filepath.Clean(cfg.StartDir + "/data/data.csv"))
	if err != nil {
		return St, err
	}

	defer func() {
		cerr := f.Close()
		if err == nil {
			err = cerr
			return
		}
	}()

	data, err := CSVToMap(f)
	if err != nil && err != io.EOF {
		return St, err
	}
	St.Logger = logger
	St.DB = data
	return St, err

}

// CSVToMap - читаем данные из CSV в мапу
func CSVToMap(reader io.Reader) (*map[string]map[string][]string, error) {
	var (
		header []string
		record []string
		err    error
	)

	r := csv.NewReader(reader)
	rows := make(map[string]map[string][]string)

	// удаляем посторонние символы из строки
	reg := regexp.MustCompile(`(?m)[+$]|,\d{1,2}\*`)
	line := 0
	for {
		record, err = r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if header == nil {
			header = record
		} else {
			var year string
			rows[year] = make(map[string][]string)
			for i := range header {
				record[i] = reg.ReplaceAllString(record[i], "")
				if header[i] == "Год/Месяц" {
					year = record[i]
					rows[year] = make(map[string][]string)
				}
				if header[i] == "Январь" {
					rows[year]["1"] = strings.Split(record[i], ",")
				}
				if header[i] == "Февраль" {
					rows[year]["2"] = strings.Split(record[i], ",")
				}
				if header[i] == "Март" {
					rows[year]["3"] = strings.Split(record[i], ",")
				}
				if header[i] == "Апрель" {
					rows[year]["4"] = strings.Split(record[i], ",")
				}
				if header[i] == "Май" {
					rows[year]["5"] = strings.Split(record[i], ",")
				}
				if header[i] == "6" {
					rows[year]["Июнь"] = strings.Split(record[i], ",")
				}
				if header[i] == "7" {
					rows[year]["Июль"] = strings.Split(record[i], ",")
				}
				if header[i] == "8" {
					rows[year]["Август"] = strings.Split(record[i], ",")
				}
				if header[i] == "Сентябрь" {
					rows[year]["9"] = strings.Split(record[i], ",")
				}
				if header[i] == "Октябрь" {
					rows[year]["10"] = strings.Split(record[i], ",")
				}
				if header[i] == "Ноябрь" {
					rows[year]["11"] = strings.Split(record[i], ",")
				}
				if header[i] == "Декабрь" {
					rows[year]["12"] = strings.Split(record[i], ",")
				}
				if header[i] == "Всего рабочих дней" {
					rows[year]["workdays"] = append(rows[year]["workdays"], record[i])
				}
				if header[i] == "Всего праздничных и выходных дней" {
					rows[year]["holidays"] = append(rows[year]["holidays"], record[i])
				}
				if header[i] == "Количество рабочих часов при 40-часовой рабочей неделе" {
					rows[year]["workhours40"] = append(rows[year]["workhours40"], record[i])
				}
				if header[i] == "Количество рабочих часов при 36-часовой рабочей неделе" {
					rows[year]["workhours36"] = append(rows[year]["workhours36"], record[i])
				}
				if header[i] == "Количество рабочих часов при 24-часовой рабочей неделе" {
					rows[year]["workhours24"] = append(rows[year]["workhours24"], record[i])
				}
			}
		}
		line++
	}
	return &rows, err
}
