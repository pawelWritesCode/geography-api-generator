package generator

import (
	"fmt"
	"testing"
)

func TestEntityGenerator_Random(t *testing.T) {
	gen := NewEntityGenerator()
	var temp Entity
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
	var temp Property
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

func TestNew(t *testing.T) {
	var e Entity
	var p Property

	randomVariable := New(e, p)

	if randomVariable.Property != p || randomVariable.Entity != e {
		t.Errorf("Initialization of RandomVariables is invalid")
	}
}

func ExampleNew() {
	e := Entity("tree")
	p := Property("height")

	randomVariable := New(e, p)
	fmt.Printf("%v", randomVariable)
	//Output:
	// {tree height}
}

func ExampleRandomVariables_EntityFU() {
	e := Entity("tree")
	p := Property("height")

	randomVariable := New(e, p)
	fmt.Printf("%s", randomVariable.EntityFU())

	//Output:
	//Tree
}

func ExampleRandomVariables_PropertyFU() {
	e := Entity("tree")
	p := Property("height")

	randomVariable := New(e, p)
	fmt.Printf("%s", randomVariable.PropertyFU())

	//Output:
	//Height
}

func TestRandomVariables_EntityFU_PropertyFU(t *testing.T) {
	type want struct {
		entityFu   string
		propertyFu string
	}
	type str struct {
		got RandomVariables
		want
	}
	data := []str{
		{
			got: RandomVariables{Entity("abc"), Property("def")},
			want: want{
				entityFu:   "Abc",
				propertyFu: "Def",
			},
		},
		{
			got: RandomVariables{Entity("1"), Property("2")},
			want: want{
				entityFu:   "1",
				propertyFu: "2",
			},
		},
		{
			got: RandomVariables{Entity("to jest zdanie"), Property("a to tez")},
			want: want{
				entityFu:   "To Jest Zdanie",
				propertyFu: "A To Tez",
			},
		},
		{
			got: RandomVariables{Entity("configProject"), Property("someProperty")},
			want: want{
				entityFu:   "ConfigProject",
				propertyFu: "SomeProperty",
			},
		},
	}

	for _, item := range data {
		if item.got.EntityFU() != item.want.entityFu {
			t.Errorf("invalid output: want %s got %s", item.want.entityFu, item.got.EntityFU())
		}

		if item.got.PropertyFU() != item.want.propertyFu {
			t.Errorf("invalid output: want %s got %s", item.want.propertyFu, item.got.PropertyFU())
		}
	}
}
