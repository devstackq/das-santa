package models

type Result struct {
	MapID       string     `json:"map_id"`
	Moves       []Children `json:"moves"`
	StackOfBags [][]int    `json:"stackOfBags"`
}
