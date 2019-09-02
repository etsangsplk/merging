package merging

type CoordinationLevelMatching struct {
	Occurances map[string]float64
}

func (m *CoordinationLevelMatching) Merge(itemsLists []Items) Items {
	seen := make(map[string]bool)
	for _, items := range itemsLists {
		for _, item := range items {
			if _, ok := m.Occurances[item.Id]; ok {
				m.Occurances[item.Id] = 1
			} else {
				m.Occurances[item.Id]++
			}
			seen[item.Id] = true
		}
	}

	items := make(map[string]Item)
	var i = 0
	for id := range seen {
		items[id] = Item{
			Id:    id,
			Score: m.Occurances[id],
		}
		i++
	}

	return flattenAndSort(items)
}
