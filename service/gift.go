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
	Sort([]models.Gift) [][]models.Gift
}

func (g *GiftSrv) Estimate() [][]models.Gift {
	log.Println("find estimate gifts", len(g.Data))

	//opt := models.NewOptimal(g.Data)
	//opt.Sort()
	//
	//vol := models.NewVolume(g.Data)
	//vol.Sort()

	log.Println("sorted by volume", len(vol.Result))
	//log.Println("sorted by weight", len(optimal.ByWeight))
	log.Println("sorted by optimal", len(opt.Result))

	return opt.Result
}
