package maps

import (
	"fmt"
	"testing"
)

func assertString(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s' want '%s' given, '%s'", got, want, "test")
	}
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertString(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := "could not find the word you were looking for"

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertString(t, err.Error(), want)
	})
}

func TestDictionary_Search(t *testing.T) {
	d := Dictionary{"test": "this is just a test"}

	got, _ := d.Search("test")
	assertString(t, got, "this is just a test")
}

func TestDictionary_Update(t *testing.T) {
	d := Dictionary{"test": "this is just a test"}

	t.Run("unknow word", func(t *testing.T) {
		_, err := d.Update("1", "2")
		want := "could not find the word you were looking for"
		assertString(t, err.Error(), want)
	})
	t.Run("normal", func(t *testing.T) {
		_, err := d.Update("test", "2")
		if err != nil {
			t.Fail()
		}
		want := "2"
		got, _ := d.Search("test")
		assertString(t, got, want)
		fmt.Println(d)
	})
}
