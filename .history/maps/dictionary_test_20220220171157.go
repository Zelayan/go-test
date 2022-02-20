package maps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {

	check := assert.New(t)
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		check.Equal(want, got)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		check.Error(err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	check := assert.New(t)

	t.Run("new world", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add("test", "this is just a test")
		want := "this is just a test"
		got, err := dictionary.Search("test")

		check.Nil(err)
		check.Equal(want, got)
	})

	t.Run("existing world", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		err := dictionary.Add("test", "new test")

		check.Error(err)

	})
}

func TestUpdate(t *testing.T) {
	check := assert.New(t)
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		newDefinition := "new definition"
		dictionary := Dictionary{word: definition}

		err := dictionary.Update(word, newDefinition)

		check.Nil(err)

		got, err := dictionary.Search(word)
		check.Nil(err)
		check.Equal(newDefinition, got)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		check.Error(err, ErrNotFound)
	})
}

func TestDelete(t *testing.T) {
	check := assert.New(t)
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{word: "test definition"}

		dictionary.Delete(word)

		_, err := dictionary.Search(word)
		check.Error(err, ErrNotFound)
	})

	t.Run("not existing word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{word: "test definition"}

		dictionary.Delete(word)

		_, err := dictionary.Search(word)
		check.Error(err, ErrNotFound)
	})

}
