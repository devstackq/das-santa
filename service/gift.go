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

func (g *GiftSrv) Estimate() models.EstimatationGifts {
	log.Println("find estimate gifts")
	//w/200 + v/100
	optimal := models.EstimatationGifts{
		Gifts:    g.Data,
		Optimal:  [][]models.Gift{},
		ByVolume: [][]models.Gift{},
		ByWeight: [][]models.Gift{},
	}

	const maxWeight = 200 //12
	const maxVolume = 100 //7

	//1 sort, min max, all items;
	//2 sort - each part gift; 7

	optimal.SortByVolumeAsc()
	optimal.SortByWeightAsc()
	optimal.SortOptimal()

	return optimal
}
