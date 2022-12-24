package models

import "sort"

type Gift struct {
	ID     int `json:"id"`
	Weight int `json:"weight"`
	Volume int `json:"volume"`
}

type EstimatationGifts struct {
	//Optimal  EstimateGift
	//ByVolume EstimateGift
	//ByWeight EstimateGift
	Gifts    []Gift
	Optimal  [][]Gift
	ByVolume [][]Gift
	ByWeight [][]Gift
}
type EstimateGift struct {
	Index      int
	Gifts      []Gift
	Min        int
	Max        int
	FirstIndex int
	LastIndex  int
}

func (eg EstimatationGifts) SortByWeightAsc() {
	temp := eg.Gifts

	sort.SliceStable(temp, func(i, j int) bool {
		return temp[i].Weight < temp[j].Weight
	})
	//eg.Gifts = data

}
func (eg EstimatationGifts) separateByVolume(data []Gift) {
	result := [][]Gift{}
	volume := 0

	var temp []Gift

	for _, gift := range data {
		//slice, sum each index item, when gift[i].Volume+=;

		if volume <= 100 {
			temp = append(temp, gift)
			volume += gift.Volume
		} else {
			result = append(result, temp)
			volume = 0
			temp = nil
		}
	}
	eg.ByVolume = result
}

func (eg EstimatationGifts) SortByVolumeAsc() {
	//eg.ByVolume = eg.Gifts
	temp := eg.Gifts

	sort.SliceStable(temp, func(i, j int) bool {
		return temp[i].Volume < temp[j].Volume
	})

	eg.separateByVolume(temp)
}

func (eg EstimatationGifts) SortOptimal() {

}
