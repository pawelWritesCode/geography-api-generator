package generators

import (
	"math/rand"
	"time"
)

type Entity string

type RandomEntity interface {
	Random() Entity
}

//RandomEntity returns randomly picked entity
func (e Entity) Random() Entity {
	entities := []Entity{
		"tree",
		"bush",
		"building",
		"sky",
		"ocean",
		"sea",
		"cloud",
		"peak",
		"land",
		"voivodeship",
		"country",
		"continent",
		"planet",
		"galactic",
		"nebula",
		"universe",
		"path",
		"star",
		"village",
		"forest",
		"reef",
	}

	rand.Seed(time.Now().UnixNano())
	min := 0
	max := len(entities) - 1
	randIndex := rand.Intn(max-min+1) + min

	return entities[randIndex]
}
