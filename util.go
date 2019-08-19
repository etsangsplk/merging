package merging

import "sort"

func flattenAndSort(seen map[string]Item) Items {
	unique := make(Items, len(seen))
	var i int
	for _, v := range seen {
		unique[i] = v
		i++
	}
	sort.Slice(unique, func(i, j int) bool {
		return unique[i].Score > unique[j].Score
	})
	return unique
}
