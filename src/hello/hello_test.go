package hello

import (
	    "testing"
)

func TestHello(t *testing.T) {
	want := "helloworld"
	if got := hello(); got != want {
		t.Errorf("hello() = %q, want %q", got, want)
	}
}

