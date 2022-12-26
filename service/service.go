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

	//s.gift = gift //set estimate gifts
	//2. find optimal path
	pathFind := NewPath(optimalGifts, data.Children, data.SnowAreas)
	pathFind.BuildGraph()
	var m = 0
	for _, element := range pathFind.sortedGifts {
		m += len(element)
	}
	log.Println("len gift", m)
	var path = pathFind.НайтиОптимальныеПути()
	log.Println("len path", len(path))
	//3. send gifts
	Result.Moves = path

	sbs := setStackOfBags(optimalGifts)
	Result.StackOfBags = sbs

	return Result, nil
}

func setStackOfBags(optimalGifts [][]models.Gift) [][]int {
	stackOfBags := [][]int{}

	for _, giftIndex := range optimalGifts {
		ids := []int{}

		for _, gift := range giftIndex {
			ids = append(ids, gift.ID)
		}
		stackOfBags = append(stackOfBags, ids)
	}
	return stackOfBags
}
