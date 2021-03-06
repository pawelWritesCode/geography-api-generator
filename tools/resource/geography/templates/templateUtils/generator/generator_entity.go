package generator

import (
	"generator/backend-go/tools/resource/geography/templates/templateUtils"
	"math/rand"
	"time"
)

type EntityGenerator struct{}

type RandomEntity interface {
	Random() templateUtils.Entity
}

//NewEntityGenerator returns new EntityGenerator struct
func NewEntityGenerator() EntityGenerator {
	return EntityGenerator{}
}

//RandomEntity returns randomly picked entity
func (e EntityGenerator) Random() templateUtils.Entity {
	entities := []templateUtils.Entity{"tree", "bush", "building", "sky", "ocean", "sea",
		"cloud", "peak", "land", "voivodeship", "country", "continent",
		"planet", "galactic", "nebula", "universe", "path", "star", "village",
		"forest", "reef",
	}

	rand.Seed(time.Now().UnixNano())
	min := 0
	max := len(entities) - 1
	randIndex := rand.Intn(max-min+1) + min

	return entities[randIndex]
}
