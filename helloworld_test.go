package main

import "testing"

func TestHello(t *testing.T) {
	assertMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}
	t.Run("say hello to somebody", func(t *testing.T) {
		got := HelloName("World")
		want := "Hello World"
		assertMessage(t, got, want)
	})
	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello()
		want := "Hello World"
		assertMessage(t, got, want)
	})

}
