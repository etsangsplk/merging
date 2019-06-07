package merging

import (
	"github.com/gonum/floats"
)

// Normaliser re-scores items by normalisation.
type Normaliser interface {
	Normalise(item Item) float64
	Init(items Items)
}

var (
	MinMaxNorm = &minMaxNorm{}
)

// minMaxNorm normalises item scores between 0 and 1 using min-max.
type minMaxNorm struct {
	min, max float64
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
	//fmt.Println(n.min, n.max)
}
