package generator

import (
	"generator/backend-go/tools/resource/geography/templates/templateUtils"
	"testing"
)

func TestEntityGenerator_Random(t *testing.T) {
	gen := NewEntityGenerator()
	var temp templateUtils.Entity
	var count int8
	for i := 0; i < 10; i++ {
		if i == 0 {
			temp = gen.Random()
			continue
		}

		if temp == gen.Random() {
			count++
		}
	}

	if count > 7 {
		t.Errorf("Entity is not generating randomly")
	}
}

func TestPropertyGenerator_Random(t *testing.T) {
	gen := NewPropertyGenerator()
	var temp templateUtils.Property
	var count int8
	for i := 0; i < 10; i++ {
		if i == 0 {
			temp = gen.Random()
			continue
		}

		if temp == gen.Random() {
			count++
		}
	}

	if count > 7 {
		t.Errorf("Property is not generating randomly")
	}
}
