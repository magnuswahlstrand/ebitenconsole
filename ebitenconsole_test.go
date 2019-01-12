package ebitenconsole

import (
	"testing"
)

func Test_Input(t *testing.T) {
	var toggle bool
	var someFloat float64
	var name string
	StringVar(&name, "name", "of a tester")
	FloatVar(&someFloat, "f", "float64 value")
	BoolVar(&toggle, "toggle", "show/hide something")

	Parse("set toggle=true")
	if toggle != true {
		t.Fatalf("expected true, got %t\n", toggle)
	}

	Parse("set toggle=false")
	if toggle != false {
		t.Fatalf("expected true, got %t\n", toggle)
	}

	Parse("set name=Testman")
	if name != "Testman" {
		t.Fatalf("expected 'Testman', got '%s'\n", name)
	}

	Parse("set f=0.15")
	expected := 0.15
	if someFloat != expected {
		t.Fatalf("expected '%f', got '%f'\n", expected, someFloat)
	}

	// Expect errors
	cmd := "set invalidcmd=Magnus"
	if err := Parse(cmd); err == nil {
		t.Fatalf("expected error from '%s'", cmd)
	}
}
