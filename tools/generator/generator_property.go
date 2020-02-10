package generator

import (
	"math/rand"
	"time"
)

//Property represents property name
type Property string

type RandomProperty interface {
	Random() Property
}

//RandomProperty returns randomly picked property
func (p Property) Random() Property {
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
