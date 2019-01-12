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

	var testVar bool
	FuncVar(func() error {
		testVar = true
		return nil
	}, "somefunc", "call a simple func")

	Parse("toggle=true")
	if toggle != true {
		t.Fatalf("expected true, got %t\n", toggle)
	}

	Parse("toggle=false")
	if toggle != false {
		t.Fatalf("expected true, got %t\n", toggle)
	}

	Parse("name=Testman")
	if name != "Testman" {
		t.Fatalf("expected 'Testman', got '%s'\n", name)
	}

	Parse("f=0.15")
	expected := 0.15
	if someFloat != expected {
		t.Fatalf("expected '%f', got '%f'\n", expected, someFloat)
	}

	//Before
	if testVar != false {
		t.Fatalf("expected 'false', got '%t'\n", testVar)
	}

	Parse("somefunc")
	//After
	if testVar != true {
		t.Fatalf("expected 'true', got '%t'\n", testVar)
	}

	// Expect errors
	cmd := "invalidvar=Magnus"
	if err := Parse(cmd); err == nil {
		t.Fatalf("expected error from '%s'", cmd)
	}

	cmd = "invalidfunc"
	if err := Parse(cmd); err == nil {
		t.Fatalf("expected error from '%s'", cmd)
	}

}
