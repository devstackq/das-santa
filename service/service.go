package service

import (
	"github.com/devstackq/das-santa.git/models"
	"log"
)

type Service struct {
	gift *GiftSrv
}

type IService interface {
	Ebash(data models.Map) error
}

func New() *Service {
	return &Service{}
}

func (s Service) Ebash(data models.Map) (models.Result, error) {
	var Result models.Result
	//1. estimate gift
	gift := NewGift(data.Gifts)
	optimalGifts := gift.Estimate()
	//var n = 0
	//for i, i2 := range optimalGifts {
	//	n += len(i2)
	//	log.Println("optimalGifts", i, len(i2))
	//}
	//log.Println("optimalGifts", n)

	//s.gift = gift //set estimate gifts
	//2. find optimal path
	pathFind := NewPath(optimalGifts, data.Children, data.SnowAreas)
	pathFind.BuildGraph()
	var path = pathFind.НайтиОптимальныеПути()
	log.Println("path", len(path))
	//3. send gifts
	Result.Moves = path

	return Result, nil
}
