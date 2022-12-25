package service

import (
	"errors"
	"github.com/devstackq/das-santa.git/models"
	"log"
	"math"
	"sort"
)

type PathFind struct {
	children    []models.Children
	sortedGifts [][]models.Gift
	snowAreas   []models.SnowArea

	SantaHome childGraphElem
	graph     []childGraphElem
}

func NewPath(sg [][]models.Gift, ch []models.Children, sa []models.SnowArea) *PathFind {
	return &PathFind{
		children:    ch,
		snowAreas:   sa,
		sortedGifts: sg,
	}
}

type wrapperGraphElement struct {
	child  *childGraphElem
	Weight float64
}

func (wge *wrapperGraphElement) CalcWeight(p2 *childGraphElem, sa models.SnowArea) (float64, error) {
	if wge.child == nil || p2 == nil {
		return 0, errors.New("nil pointer")
	}
	if wge.child.X == p2.X && wge.child.Y == p2.Y {
		return 0, nil
	}
	//todo if point in snow area
	var minSpeed float64 = 10 //km/h
	var maxSpeed float64 = 70 //km/h
	var distance = math.Sqrt(math.Pow(float64(wge.child.X-p2.X), 2) + math.Pow(float64(wge.child.Y-p2.Y), 2))
	var k = float64(wge.child.Y-p2.Y) / float64(wge.child.X-p2.X)
	var n = float64(p2.Y) - k*float64(p2.X) - float64(sa.Y)
	var D = 4*math.Pow(n*k-float64(sa.X), 2) - 4*(1+math.Pow(k, 2))*(math.Pow(n, 2)+math.Pow(float64(sa.X), 2)-math.Pow(float64(sa.Radius), 2))
	if D <= 0 {
		wge.Weight = distance / maxSpeed
		return distance / maxSpeed, nil
	}
	var x1 = (-2*n*k + 2*float64(sa.X) + math.Sqrt(D)) / (2 * (1 + math.Pow(k, 2)))
	var x2 = (-2*n*k + 2*float64(sa.X) - math.Sqrt(D)) / (2 * (1 + math.Pow(k, 2)))
	var y1 = k*(x1-float64(p2.X)) + float64(p2.Y)
	var y2 = k*(x2-float64(p2.X)) + float64(p2.Y)

	// distance to snow area
	var distanceToSnowArea = math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))

	wge.Weight = distanceToSnowArea/minSpeed + (distance-distanceToSnowArea)/maxSpeed
	return distanceToSnowArea/minSpeed + (distance-distanceToSnowArea)/maxSpeed, nil
}

type childGraphElem struct {
	X       int
	Y       int
	toHome  float64
	donated bool
	next    []wrapperGraphElement
}

func (pf *PathFind) BuildGraph() {
	//1. build graph
	pf.SantaHome = childGraphElem{
		X: 0,
		Y: 0,
	}
	for _, child := range pf.children {
		pf.graph = append(pf.graph, childGraphElem{
			X: child.X,
			Y: child.Y,
		})
	}
	for i := range pf.graph {
		var vertex = wrapperGraphElement{
			child:  &pf.graph[i],
			Weight: 0,
		}
		var maxlen float64 = 0
		for _, area := range pf.snowAreas {
			var weight, err = vertex.CalcWeight(&pf.SantaHome, area)
			if err != nil {
				log.Println(err)
				continue
			}
			if weight > maxlen {
				maxlen = weight
			}
		}
		vertex.Weight = maxlen
		pf.SantaHome.next = append(pf.SantaHome.next, vertex)
	}
	for i := range pf.graph {
		pf.graph[i].toHome = pf.SantaHome.next[i].Weight
		for j := range pf.graph {
			if i == j {
				continue
			}
			var vertex = wrapperGraphElement{
				child:  &pf.graph[j],
				Weight: 0,
			}
			var maxlen float64 = 0
			for _, area := range pf.snowAreas {
				var weight, err = vertex.CalcWeight(&pf.graph[i], area)
				if err != nil {
					log.Println(err)
					continue
				}
				if weight > maxlen {
					maxlen = weight
				}
			}
			vertex.Weight = maxlen
			pf.graph[i].next = append(pf.graph[i].next, vertex)
		}
		sort.SliceStable(pf.graph[i].next, func(k, j int) bool {
			return pf.graph[i].next[k].Weight < pf.graph[i].next[j].Weight
		})
	}
	sort.SliceStable(pf.SantaHome.next, func(i, j int) bool {
		return pf.SantaHome.next[i].Weight < pf.SantaHome.next[j].Weight
	})
}
func (pf *PathFind) НайтиОптимальныеПути() []models.Children {
	sort.SliceStable(pf.sortedGifts, func(i, j int) bool {
		return len(pf.sortedGifts[i]) > len(pf.sortedGifts[j])
	})
	var path []models.Children
	for i := len(pf.sortedGifts) - 1; i > 0; i-- {
		var N = len(pf.sortedGifts[i])
		var now = &pf.SantaHome
		for j := 0; j < N-1; j++ {
			for k, element := range now.next {
				if !element.child.donated {
					now.next[k].child.donated = true
					path = append(path, models.Children{
						X: element.child.X,
						Y: element.child.Y,
					})
					now = element.child
					break
				}
			}
		}
		path = append(path, models.Children{
			X: 0,
			Y: 0,
		})
	}
	return path
}
