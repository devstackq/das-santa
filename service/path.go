package service

import "github.com/devstackq/das-santa.git/models"

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

}
func (pf *PathFind) РаздатьПодарки() {}
