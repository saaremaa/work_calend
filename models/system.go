package models

// MessageUnauthorized - свое сообщение при событии "пользователь не авторизован"
type MessageUnauthorized struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// MsgStat - структура ответа на запрос статистических данных
type MsgStat struct {
	Error      string `json:"error"`
	StartYear  string `json:"start_year"`
	StopYear   string `json:"stop_year"`
	LastUpdate string `json:"last_update"`
	UPTime     string `json:"uptime"`
}
