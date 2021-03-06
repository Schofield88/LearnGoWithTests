package main

import "testing"

func TestHello (t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("Received: %q but Expected: %q", got, want)
		}
	}

	t.Run("Greets a person by name", func(t *testing.T) {
		got := Hello("Luke", "")
		want := "Hello, Luke"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Defaults to Hello, World", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("French greeting", func(t *testing.T) {
		got := Hello("Marie", "French")
		want := "Salut, Marie"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Czech greeting", func(t *testing.T) {
		got := Hello("Marta", "Czech")
		want := "Cao, Marta"

		assertCorrectMessage(t, got, want)
	})
}
