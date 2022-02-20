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
	t.Run("existing word", func(t *testing.T) {
		
	})
}
