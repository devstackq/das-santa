package models

type Result struct {
	MapID       string     `json:"mapID"`
	Moves       []Children `json:"moves"`
	StackOfBags [][]int    `json:"stackOfBags"`
}
