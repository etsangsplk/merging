package merging

import (
	"github.com/hscells/trecresults"
	"math"
)

// Item is an element in a list.
// It has an Id which should be unique to the list it appears.
// It has a Score which indicates the rank position of the item in a list.
type Item struct {
	Id    string
	Score float64
}

// Items is a list of items.
type Items []Item

// Merger merges multiple lists into a single list.
type Merger interface {
	Merge(itemsLists []Items) Items
}

// Sort sorts items in a list by their score.
//func (it Items) Sort() {
//	sort.Slice(it, func(i, j int) bool {
//		return it[i].Score > it[j].Score
//	})
//}

func (it Items) TRECResults(topic string) trecresults.ResultList {
	var list trecresults.ResultList
	for i, item := range it {
		list = append(list, &trecresults.Result{
			Topic:     topic,
			Iteration: "0",
			DocId:     item.Id,
			Rank:      int64(i) + 1,
			Score:     item.Score,
			RunName:   topic,
		})
	}
	return list
}

func (it Items) Max() float64 {
	var max float64
	for _, item := range it {
		if item.Score > max {
			max = item.Score
		}
	}
	return max
}

func (it Items) Min() float64 {
	min := math.MaxFloat64
	for _, item := range it {
		if item.Score < min {
			min = item.Score
		}
	}
	if min == math.MaxFloat64 {
		return 0
	}
	return min
}

func (it Items) Cut(kappa float64) Items {
	var items Items
	for _, item := range it {
		if item.Score >= kappa {
			items = append(items, item)
		}
	}
	return items
}

func (it Items) Sum() float64 {
	var sum float64
	for _, item := range it {
		sum += item.Score
	}
	return sum
}

func FromTRECResults(list trecresults.ResultList) Items {
	items := make(Items, len(list))
	for i, v := range list {
		items[i] = Item{
			Id:    v.DocId,
			Score: v.Score,
		}
	}
	return items
}
