package hello

import "testing"

func TestHello(t *testing.T) {
	expected := "Hello, world."
	result := Hello()
	if result != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, result)
	}
}
