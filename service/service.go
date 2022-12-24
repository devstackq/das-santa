package service

import "github.com/devstackq/das-santa.git/models"

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
	pathFind.НайтиОптимальныеПути()
	pathFind.РаздатьПодарки()

	//3. send gifts

	return Result, nil
}
