package models

type Result struct {
	MapID       string     `json:"mapID"`
	Moves       []Children `json:"moves"`
	StackOfBags [][]int    `json:"stackOfBags"`
}

type ResponseSendRound struct {
	RoundID string `json:"roundId"`
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

type ResponseGetRound struct {
	Success bool `json:"success"`
	Data    `json:"data"`
}
type Data struct {
	ErrorMessage string `json:"error_message"`
	Status       string `json:"status"`
	TotalTime    int    `json:"total_time"`
	TotalLength  int    `json:"total_length"`
}
