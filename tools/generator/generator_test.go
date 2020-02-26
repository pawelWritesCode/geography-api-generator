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

func TestEntity_EntityFU(t *testing.T) {
	e := Entity("hello")
	want := "Hello"

	if e.EntityFU() != want {
		t.Errorf("got %s want %s", e.EntityFU(), want)
	}
}

func ExampleEntity_EntityFU() {
	e := Entity("hello")
	fmt.Println(e.EntityFU())

	//Output:
	//Hello
}

func TestProperty_PropertyFU(t *testing.T) {
	p := Property("width")
	want := "Width"

	if p.PropertyFU() != want {
		t.Errorf("got %s want %s", p.PropertyFU(), want)
	}
}

func ExampleProperty_PropertyFU() {
	p := Property("width")
	fmt.Println(p.PropertyFU())

	//Output:
	//Width
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
		if item.got.Entity.EntityFU() != item.want.entityFu {
			t.Errorf("invalid output: want %s got %s", item.want.entityFu, item.got.Entity.EntityFU())
		}

		if item.got.Property.PropertyFU() != item.want.propertyFu {
			t.Errorf("invalid output: want %s got %s", item.want.propertyFu, item.got.Property.PropertyFU())
		}
	}
}
