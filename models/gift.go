package models

import (
	"math"
	"sort"
)

type Gift struct {
	ID     int     `json:"id"`
	Weight float64 `json:"weight"`
	Volume float64 `json:"volume"`
}

type EstimatationGifts struct {
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

const maxWeight = 200.0 //12
const maxVolume = 100.0 //7

type Weight struct {
}
type Volume struct {
}
type Optimal struct {
	gifts  []Gift
	Result [][]Gift
}

func NewOptimal(data []Gift) *Optimal {
	return &Optimal{
		gifts: data,
	}
}

func (o *Optimal) Sort() {
	type IndexValue struct {
		IndexGift int
		Value     float64
	}
	temp := o.gifts
	var opts []IndexValue

	//prepare data
	for idx, gift := range temp {
		diff := (gift.Weight / maxWeight) - (gift.Volume / maxVolume)
		opt := IndexValue{}
		opt.IndexGift = idx
		opt.Value = math.Abs(diff)

		opts = append(opts, opt)
	}

	sort.SliceStable(opts, func(i, j int) bool {
		return opts[i].Value < opts[j].Value
	})

	sorted := make([]Gift, len(opts))

	for idx, item := range opts {
		sorted[idx] = temp[item.IndexGift]
	}

	o.separate(sorted)

}
func (o *Optimal) separate(sorted []Gift) {

	var result [][]Gift

	sumVolume := 0.0
	sumWeight := 0.0

	var temp []Gift

	for _, gift := range sorted {
		if sumWeight <= maxWeight && sumWeight+gift.Weight <= maxWeight || sumVolume <= maxVolume && sumVolume+gift.Volume <= maxVolume {
			temp = append(temp, gift)
			sumVolume += gift.Volume
			sumWeight += gift.Weight
		} else {
			result = append(result, temp)
			sumVolume = 0
			sumWeight = 0
			temp = nil
		}
	}

	o.Result = result
}

func (eg *EstimatationGifts) SortByWeightAsc() {
	temp := eg.Gifts

	sort.SliceStable(temp, func(i, j int) bool {
		return temp[i].Weight < temp[j].Weight
	})
	eg.separateByWeight(temp)

}

func (eg *EstimatationGifts) separateByWeight(data []Gift) {
	result := [][]Gift{}
	weight := 0.0

	var temp []Gift

	for _, gift := range data {
		//slice, sum each index item, when gift[i].Volume+=;
		if weight <= maxWeight && weight+gift.Weight <= maxWeight {
			temp = append(temp, gift)
			weight += gift.Weight
		} else {
			result = append(result, temp)
			weight = 0
			temp = nil
		}
	}

	eg.ByWeight = result
}
func (eg *EstimatationGifts) separateByVolume(data []Gift) {
	result := [][]Gift{}
	volume := 0.0

	var temp []Gift

	for _, gift := range data {
		//slice, sum each index item, when gift[i].Volume+=;
		if volume <= maxVolume && volume+gift.Volume <= maxVolume {
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

func (eg *EstimatationGifts) SortByVolumeAsc() {
	//eg.ByVolume = eg.Gifts
	temp := eg.Gifts

	sort.SliceStable(temp, func(i, j int) bool {
		return temp[i].Volume < temp[j].Volume
	})

	eg.separateByVolume(temp)
}
