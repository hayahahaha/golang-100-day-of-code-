package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("test search know world", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertSearch(t, got, want)
	})

	t.Run("test search unknow world", func(t *testing.T) {
		_, err := dictionary.Search("test1")
		want := "not found"

		asseartError(t, err)
		assertSearch(t, err.Error(), want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("testing add new world", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add("test", "this is just a test")
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertSearch(t, got, want)
	})
}

func asseartError(t testing.TB, err error) {
	if err == nil {
		t.Fatal("expected get an error")
	}
}

func assertSearch(t testing.TB, got, want string) {
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
