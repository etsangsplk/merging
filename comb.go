package merging

// CombSUM sums an item's score from all lists where it was present.
type CombSUM struct {
	Normaliser
}

func (c CombSUM) Merge(itemsLists []Items) Items {
	var (
		unique Items
		i      int
	)

	// Sum item scores from all the item lists where it is present.
	seen := make(map[string]Item)
	for _, items := range itemsLists {
		for _, item := range items {
			score := c.Normaliser(item, items)
			if _, ok := seen[item.Id]; !ok {
				item.Score = score
				seen[item.Id] = item
			} else {
				item.Score++
			}
		}
	}

	// Create a flat slice from the unique items.
	unique = make(Items, len(seen))
	for _, v := range seen {
		unique[i] = v
		i++
	}

	return unique
}

// CombMNZ additionally multiplies the CombSUM score by the number of lists that contain that item.
type CombMNZ struct {
	Normaliser
}

func (c CombMNZ) Merge(itemsLists []Items) Items {
	// Compute the CombSUM score for each item.
	csum := CombSUM{Normaliser: c.Normaliser}
	its := csum.Merge(itemsLists)

	// Then, record how many times each item appears in each of the lists of items.
	k := make(map[string]float64)
	for _, items := range itemsLists {
		for _, item := range items {
			if _, ok := k[item.Id]; !ok {
				k[item.Id] = 1
			} else {
				k[item.Id]++
			}
		}
	}

	// Finally, multiply each item score by the number of times a list of items contained that item.
	for _, item := range its {
		item.Score *= k[item.Id]
	}

	return its
}
