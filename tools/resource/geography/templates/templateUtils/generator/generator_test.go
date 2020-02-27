package generator

import (
	"fmt"
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

func TestNew(t *testing.T) {
	var e templateUtils.Entity
	var p templateUtils.Property

	randomVariable := templateUtils.NewTemplateVariables(e, p)

	if randomVariable.Property != p || randomVariable.Entity != e {
		t.Errorf("Initialization of TemplateVariables is invalid")
	}
}

func ExampleNew() {
	e := templateUtils.Entity("tree")
	p := templateUtils.Property("height")

	randomVariable := templateUtils.NewTemplateVariables(e, p)
	fmt.Printf("%v", randomVariable)
	//Output:
	// {tree height}
}

func TestEntity_EntityFU(t *testing.T) {
	e := templateUtils.Entity("hello")
	want := "Hello"

	if e.EntityFU() != want {
		t.Errorf("got %s want %s", e.EntityFU(), want)
	}
}

func ExampleEntity_EntityFU() {
	e := templateUtils.Entity("hello")
	fmt.Println(e.EntityFU())

	//Output:
	//Hello
}

func TestProperty_PropertyFU(t *testing.T) {
	p := templateUtils.Property("width")
	want := "Width"

	if p.PropertyFU() != want {
		t.Errorf("got %s want %s", p.PropertyFU(), want)
	}
}

func ExampleProperty_PropertyFU() {
	p := templateUtils.Property("width")
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
		got templateUtils.TemplateVariables
		want
	}
	data := []str{
		{
			got: templateUtils.TemplateVariables{templateUtils.Entity("abc"), templateUtils.Property("def")},
			want: want{
				entityFu:   "Abc",
				propertyFu: "Def",
			},
		},
		{
			got: templateUtils.TemplateVariables{templateUtils.Entity("1"), templateUtils.Property("2")},
			want: want{
				entityFu:   "1",
				propertyFu: "2",
			},
		},
		{
			got: templateUtils.TemplateVariables{templateUtils.Entity("to jest zdanie"), templateUtils.Property("a to tez")},
			want: want{
				entityFu:   "To Jest Zdanie",
				propertyFu: "A To Tez",
			},
		},
		{
			got: templateUtils.TemplateVariables{templateUtils.Entity("configProject"), templateUtils.Property("someProperty")},
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
