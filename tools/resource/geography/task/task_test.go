package task

import "testing"

func TestNew(t *testing.T) {
	got := New()
	want := Task{}

	if got != want {
		t.Errorf("Instantiating new task failed")
	}
}
