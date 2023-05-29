package maps

import "testing"

func TestSearch(t *testing.T) {

	assertString := func(t testing.TB, got string, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %s, but want %s", got, want)
		}
	}

	assertError := func(t testing.TB, got, want error) {
		t.Helper()

		if got != want {
			t.Errorf("got error %q want %q", got, want)
		}
	}

	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("test dictionary for word test", func(t *testing.T) {

		got, err := dictionary.Search("test")
		want := "this is just a test"

		if err != nil {
			t.Errorf("got error %q want %q", got, want)
		}

		assertString(t, got, want)
	})

	t.Run("test not in a dictionary", func(t *testing.T) {

		_, got := dictionary.Search("unknown")

		assertError(t, got, ErrNotFound)
	})

	t.Run("test add word to dictionary", func(t *testing.T) {
		dict := Dictionary{}

		word := "dog"
		description := "peoples best friend"
		dict.Add(word, description)

		got, err := dict.Search("dog")

		if err != nil {
			t.Fatal("should find added word")
		}

		assertString(t, got, description)
	})
}
