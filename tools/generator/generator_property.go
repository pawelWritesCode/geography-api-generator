package generator

import (
	"math/rand"
	"time"
)

//Property represents property name
type Property string

//PropertyGenerator represents object that generates random properties
type PropertyGenerator struct{}

type RandomProperty interface {
	Random() Property
}

//NewPropertyGenerator returns new PropertyGenerator struct
func NewPropertyGenerator() PropertyGenerator {
	return PropertyGenerator{}
}

//RandomProperty returns randomly picked property
func (p PropertyGenerator) Random() Property {
	properties := []Property{
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
