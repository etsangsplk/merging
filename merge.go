package merging

import (
	"github.com/hscells/trecresults"
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
