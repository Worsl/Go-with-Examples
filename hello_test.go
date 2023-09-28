package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestBye(t *testing.T) {
	got := Bye()
	want := "Bye, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
