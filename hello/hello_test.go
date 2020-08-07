package hello

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	if got != "Hello, World!" {
		t.Errorf("Hello() = %q, want Hello, World!", got)
	}
}
