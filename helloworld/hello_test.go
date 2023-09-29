package main

import "testing"


func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris","")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("","")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T){
		got := Hello("Benzema", "French")
		want := "Bonjour, Benzema"
		assertCorrectMessage(t,got,want)

	})
}



func assertCorrectMessage(t testing.TB, got , want string){
	t.Helper() // Hey, testing system, if something goes wrong in this helper function, show it as if it happened in the line where I called the helper." This makes it easier to spot and fix issues in your test code because the error message will point to the exact place in your test where the problem occurred.
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}