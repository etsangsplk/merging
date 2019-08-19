package merging

import (
	"github.com/gonum/floats"
	"github.com/montanaflynn/stats"
	"gonum.org/v1/gonum/stat"
)

// Normaliser re-scores items by normalisation.
type Normaliser interface {
	Normalise(item Item) float64
	Init(items Items)
}

var (
	MinMaxNorm  = &minMaxNorm{}
	ZScoreNorm  = &zScoreNorm{}
	SoftmaxNorm = &softmaxNorm{}
)

// minMaxNorm normalises item scores between 0 and 1 using min-max.
type minMaxNorm struct {
	min, max float64
}

type zScoreNorm struct {
	mean, std float64
}

type softmaxNorm struct {
	items map[string]float64
}

func (n softmaxNorm) Normalise(item Item) float64 {
	return n.items[item.Id]
}

func (n softmaxNorm) Init(items Items) {
	if len(items) == 0 {
		return
	}
	n.items = make(map[string]float64)
	s := make([]float64, len(items))
	for i, item := range items {
		s[i] = item.Score
	}
	norm, err := stats.SoftMax(s)
	if err != nil {
		panic(err)
	}
	for i, item := range items {
		n.items[item.Id] = norm[i]
	}
}

func (n zScoreNorm) Normalise(item Item) float64 {
	return (item.Score - n.mean) / n.std
}

func (n zScoreNorm) Init(items Items) {
	if len(items) == 0 {
		return
	}
	s := make([]float64, len(items))
	for i, item := range items {
		s[i] = item.Score
	}
	n.mean = stat.Mean(s, nil)
	n.std = stat.StdDev(s, nil)
}

func (n *minMaxNorm) Normalise(item Item) float64 {
	if n.max == n.min {
		return 0.0
	}
	return (item.Score - n.min) / (n.max - n.min)
}

// (18.0907 - 0) / (19.49492965207815 - 0)

func (n *minMaxNorm) Init(items Items) {
	if len(items) == 0 {
		n.min = 0
		n.max = 0
		return
	}
	s := make([]float64, len(items))
	for i, item := range items {
		s[i] = item.Score
	}
	n.min = floats.Min(s)
	n.max = floats.Max(s)
	//fmt.Println("min, max:", n.min, n.max)
}

func Normalise(n Normaliser, items Items) Items {
	n.Init(items)
	normalised := make(Items, len(items))
	for i, item := range items {
		normalised[i] = Item{
			Id:    item.Id,
			Score: n.Normalise(item),
		}

	}
	return normalised
}
