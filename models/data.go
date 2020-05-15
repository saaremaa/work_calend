package models

// MsgIsWorkDay - структура ответа на запрос является ли день рабочим
type MsgIsWorkDay struct {
	Error   string `json:"error"`
	Date    string `json:"date"`
	WorkDay bool   `json:"workday"`
}
