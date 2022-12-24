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

func (g *GiftSrv) Estimate() {
	log.Println("find estimate gifts")

}
