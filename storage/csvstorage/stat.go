package csvstorage

import (
	"sort"
)

// Statistic cтруктура для статистических данных
type Statistic struct {
	MinYear    string
	MaxYear    string
	LastUpdate string
}

// CaclStat - считаем статистические параметры датасета
func (db *CSVStorage) CaclStat() (stat Statistic) {
	rows := *db.DB
	keys := make([]string, 0, len(rows))
	for k := range rows {
		if k != "" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	stat.MinYear = keys[0]
	stat.MaxYear = keys[len(keys)-1]
	return
}
