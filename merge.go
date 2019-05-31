package merging

import "sort"

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
func (it Items) Sort() {
	sort.Slice(it, func(i, j int) bool {
		return it[i].Score > it[j].Score
	})
}
