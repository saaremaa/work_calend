package csvstorage

import (
	"strconv"
	"time"
)

// IsHoliday проверяем что день является праздничным
func (db *CSVStorage) IsHoliday(date string) (bool, error) {
	d, err := time.Parse("02.01.2006", date)
	if err != nil {
		return false, err
	}
	year := strconv.Itoa(d.Year())
	month := strconv.Itoa(int(d.Month()))
	day := strconv.Itoa(d.Day())
	data := *db.DB
	isHoliday := false
	for i := range data[year][month] {
		if day == data[year][month][i] {
			isHoliday = true
			break
		}
	}
	return isHoliday, nil
}
