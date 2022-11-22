package main

import "testing"

func assert(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}

func TestHello(t *testing.T) {
	t.Run("Saying Hello to ppl", func(t *testing.T) {
		got := Hello("kamuri", "french")
		want := "Bonjour kamuri"
		assert(t, got, want)
	})
	t.Run("Defaults to bonjour in case of empty str", func(t *testing.T) {
		got := Hello("", "french")
		want := "Bonjour"
		assert(t, got, want)
	})
	t.Run("in portuguese", func(t *testing.T) {
		got := Hello("kamuri", "portuguese")
		want := "Bom dia kamuri"
		assert(t, got, want)
	})
	t.Run("in indonesian", func(t *testing.T) {
		got := Hello("kamuri", "indonesian")
		want := "Selamat pagi kamuri"
		assert(t, got, want)
	})
}
