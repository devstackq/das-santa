package service

import (
	"github.com/devstackq/das-santa.git/models"
	"log"
)

type GiftSrv struct {
	Data []models.Gift
}

func NewGift(gifts []models.Gift) *GiftSrv {
	return &GiftSrv{
		Data: gifts,
	}
}

type Sortir interface {
	Sort([]models.Gift)
}

func (g *GiftSrv) Estimate() [][]models.Gift {
	log.Println("find estimate gifts", len(g.Data))

	s := models.EstimatationGifts{
		Gifts:   g.Data,
		Optimal: [][]models.Gift{},
	}
	opt := models.NewOptimal(g.Data)
	opt.Sort()
	//vol := models.NewVolume(g.Data)
	//vol.Sort()

	s.SortByVolumeAsc()
	s.SortByWeightAsc()

	log.Println("sorted by volume", len(s.ByVolume))
	log.Println("sorted by opt", len(opt.Result), len(opt.Result))
	log.Println("sorted by wei", len(s.ByWeight))

	return opt.Result //FIXME try to other sorted data
}
