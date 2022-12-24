package models

type Map struct {
	Gifts     []Gift     `json:"gifts"`
	SnowAreas []SnowArea `json:"snowAreas"`
	Children  []Children `json:"children"`
}
