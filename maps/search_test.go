package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("test search know world", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})
	t.Run("test search unknow world", func(t *testing.T) {
		_, err := dictionary.Search("test1")
		want := "not found"

		// assertError(t, err, nil)
		assertStrings(t, err.Error(), want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("testing add new world", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"
		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})
	t.Run("testing add new existed world", func(t *testing.T) {
		word := "test"
		definition := "this first definition"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "Second definition")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update success", func(t *testing.T) {
		word := "test"
		definition := "this just a test"
		newDefinition := "this just a test 2"
		dictionary := Dictionary{word: definition}
		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})
	t.Run("Update fail", func(t *testing.T) {

	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete success", func(t *testing.T) {
		word := "test"
		definition := "this just a test"
		dictionary := Dictionary{word: definition}
		dictionary.Delete(word)

		_, err := dictionary.Search(word)

		if err != ErrNotFound {
			t.Errorf("expect %q to be delete: ", word)
		}

	})
	t.Run("Delete success word not exist", func(t *testing.T) {

	})
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added world")
	}
	assertStrings(t, got, definition)
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertStrings(t testing.TB, got, want string) {
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
