package generators

import (
	"math/rand"
	"time"
)

//Property represents property name
type Property string

//RandomProperty returns randomly picked property
func RandomProperty() Property  {
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
	randIndex := rand.Intn(max - min + 1) + min

	return properties[randIndex]
}