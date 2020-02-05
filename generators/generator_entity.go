package generators

import (
	"math/rand"
	"time"
)

type Entity string

//RandomEntity returns randomly picked entity
func RandomEntity() Entity  {
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
	randIndex := rand.Intn(max - min + 1) + min

	return entities[randIndex]
}