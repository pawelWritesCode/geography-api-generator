package generator

import (
	"generator/backend-go/tools/resource/geography/templates/templateUtils"
	"math/rand"
	"time"
)

//PropertyGenerator represents object that generates random properties
type PropertyGenerator struct{}

type RandomProperty interface {
	Random() templateUtils.Property
}

//NewPropertyGenerator returns new PropertyGenerator struct
func NewPropertyGenerator() PropertyGenerator {
	return PropertyGenerator{}
}

//RandomProperty returns randomly picked property
func (p PropertyGenerator) Random() templateUtils.Property {
	properties := []templateUtils.Property{
		"height",
		"width",
		"length",
		"opacity",
		"volume",
		"field",
		"area",
	}

	rand.Seed(time.Now().UnixNano())
	min := 0
	max := len(properties) - 1
	randIndex := rand.Intn(max-min+1) + min

	return properties[randIndex]
}
