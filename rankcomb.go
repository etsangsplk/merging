package merging

type Borda struct{}

func (n Borda) Merge(itemsLists []Items) Items {
	// Sum item scores from all the item lists where it is present.
	seen := make(map[string]Item)
	for _, items := range itemsLists {
		n := float64(len(items))
		for i, item := range items {
			seen[item.Id] = Item{
				Id:    item.Id,
				Score: seen[item.Id].Score + ((n - float64(i)) / n),
			}
		}
	}

	return flattenAndSort(seen)
}
