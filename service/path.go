package service

import (
	"github.com/devstackq/das-santa.git/models"
	"log"
)

type PathFind struct {
	children    []models.Children
	sortedGifts [][]models.Gift
	snowAreas   []models.SnowArea
}

func NewPath(sg [][]models.Gift, ch []models.Children, sa []models.SnowArea) *PathFind {
	return &PathFind{
		children:    ch,
		snowAreas:   sa,
		sortedGifts: sg,
	}
}

func (pf *PathFind) НайтиОптимальныеПути() {
	log.Print(pf.children, pf.snowAreas, pf.sortedGifts, "starting find")
}
func (pf *PathFind) РаздатьПодарки() {}
