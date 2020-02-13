package templates

import (
	"generator/backend-go/tools/generator"
	"generator/backend-go/tools/resource"
	"testing"
	"text/template"
)

func TestNew(t *testing.T) {
	res := resource.New("abc", "def")
	tpl := template.New("abc")
	variable := generator.New(generator.Entity("a"), generator.Property("b"))
	got := New(res, tpl, variable)

	if got.Resource != res {
		t.Errorf("invalid resource got %v, has %v", got.Resource, res)
	}

	if got.Payload != tpl {
		t.Errorf("invalid payload got %v, has %v", got.Payload, tpl)
	}

	if got.Variables != variable {
		t.Errorf("invalid variables got %v, has %v", got.Variables, variable)
	}
}
