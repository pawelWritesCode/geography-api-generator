package resource

import "testing"

func TestNew(t *testing.T) {
	dir := "a"
	fileName := "b"

	resource := New(dir, fileName)

	if resource.Directory != dir || resource.FileName != fileName {
		t.Errorf("Initialization of Resource is invalid")
	}
}
