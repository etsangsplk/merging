package merging

import "math"

// Normaliser re-scores items by normalisation.
type Normaliser func(item Item, items Items) float64

// Normalise item scores between 0 and 1 using min-max.
func MinMaxNorm(item Item, items Items) float64 {
	var (
		min, max float64
	)
	min = math.MaxFloat64
	for _, i := range items {
		if i.Score > max {
			max = i.Score
		}
		if i.Score < min {
			min = i.Score
		}
	}

	return (item.Score - min) / (max - min)
}
