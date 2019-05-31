package merging

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

var (
	itemsLists = []Items{
		{
			{Id: "1", Score: 10},
			{Id: "2", Score: 9},
			{Id: "3", Score: 8},
			{Id: "4", Score: 7},
			{Id: "5", Score: 6},
			{Id: "6", Score: 5},
			{Id: "7", Score: 4},
		},
		{
			{Id: "1", Score: 542},
			{Id: "2", Score: 343},
			{Id: "4", Score: 243},
			{Id: "5", Score: 211},
			{Id: "6", Score: 135},
			{Id: "7", Score: 122},
			{Id: "8", Score: 111},
			{Id: "9", Score: 80},
			{Id: "10", Score: 40},
			{Id: "3", Score: 10},
		},
		{
			{Id: "1", Score: 76},
			{Id: "12", Score: 12},
			{Id: "2", Score: 9},
			{Id: "4", Score: 8},
			{Id: "5", Score: 4},
			{Id: "6", Score: 2},
			{Id: "7", Score: 1},
			{Id: "8", Score: 0},
			{Id: "3", Score: -1},
			{Id: "9", Score: -12},
		},
	}
)

func TestCombSUM_MinMaxNorm(t *testing.T) {
	c := CombSUM{Normaliser: MinMaxNorm}
	it := c.Merge(itemsLists)
	it.Sort()
	ids := []string{"1", "2", "3", "4", "5", "12", "8", "6", "9", "10", "7"}
	its := make([]string, len(ids))
	for i, x := range it {
		its[i] = x.Id
	}
	assert.Equal(t, its, ids)
}

func TestCombMNZ_MinMaxNorm(t *testing.T) {
	c := CombMNZ{Normaliser: MinMaxNorm}
	it := c.Merge(itemsLists)
	it.Sort()
	ids := []string{"1", "2", "3", "4", "5", "12", "8", "6", "9", "10", "7"}
	its := make([]string, len(ids))
	for i, x := range it {
		its[i] = x.Id
	}
	assert.Equal(t, its, ids)
}
