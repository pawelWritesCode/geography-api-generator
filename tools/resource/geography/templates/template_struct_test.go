package templates

import (
	"generator/backend-go/tools/resource"
	"generator/backend-go/tools/resource/geography/templates/templateUtils"
	"testing"
	"text/template"
)

func TestNew(t *testing.T) {
	res := resource.New("abc", "def")
	tpl := template.New("abc")
	variable := templateUtils.NewTemplateVariables(templateUtils.Entity("a"), templateUtils.Property("b"))
	got := New(res, tpl, variable)

	if Resource != res {
		t.Errorf("invalid resource got %v, has %v", Resource, res)
	}

	if got.Payload != tpl {
		t.Errorf("invalid payload got %v, has %v", got.Payload, tpl)
	}

	if got.Variables != variable {
		t.Errorf("invalid variables got %v, has %v", got.Variables, variable)
	}
}
